---
date: '2007-12-25T12:00:00-00:00'
language: en
tags:
- bbcode
- i18n
- phpbb
- phpbb3
title: Translating Custom BBCode Buttons
---


Recently, in a support request on [phpbb.com](http://www.phpbb.com) the question came up, if it's also possible, in the
situation that you have a multi-lingual board, that custom BBCode-buttons
could also have a localized button label and tooltip. While phpBB3 by itself
doesn't really support this out of the box, it's not all that hard to
get it working.

-------------------------------


**Note:** Use the code presented here at your own risk. It was only shortly 
tested on my local devboard. Also don't forget to backup your files
before applying these changes.

The core to this whole task is the ``display_custom_bbcode()`` function inside
the {$phpbb\_root\_path}/includes/functions_display.php since this is where
phpBB3 fetches the custom BBCode from the database and generates the 
respective template variables. 

Of primary interest within this function is following section (L823 - L834)
    
    while ($row = $db->sql_fetchrow($result))
	{
		$template->assign_block_vars('custom_tags', array(
			'BBCODE_NAME'		=> "'[{$row['bbcode_tag']}]', '[/" . str_replace('=', '', $row['bbcode_tag']) . "]'",
			'BBCODE_ID'			=> $num_predefined_bbcodes + ($i * 2),
			'BBCODE_TAG'		=> $row['bbcode_tag'],
			'BBCODE_HELPLINE'	=> $row['bbcode_helpline'],
			'A_BBCODE_HELPLINE'	=> str_replace(array('&amp;', '&quot;', "'", '&lt;', '&gt;'), array('&', '"', "\'", '<', '>'), $row['bbcode_helpline']),
		));

		$i++;
	}
	
This is the place were we want to get the multilanguage support in. Here
we will have to replace the values for ``BBCODE_TAG``, ``BBCODE_HELPLINE`` and 
``A_BBCODE_HELPLINE``. The idea is to take entries from the {$phpbb\_root\_path}/language/\*/posting.php for these fields and fall back
to the original value (as set in the ACP) if it isn't present in the 
currently selected language pack.

The resulting code looks like this:
    
    while ($row = $db->sql_fetchrow($result))
	{
	    // Fetch a translation for the custom bbcode if it exists
	    $helpline = $row['bbcode_helpline'];
	    $helpline_idx = 'CUSTOM_BBCODE_HELP_' . $row['bbcode_tag'];
	    $tag_idx = 'CUSTOM_BBCODE_TAG_' . $row['bbcode_tag'];
	    
	    if (isset($user->lang[$helpline_idx]))
	    {
	        $helpline = $user->lang[$helpline_idx];
	    }
	    if (isset($user->lang[$tag_idx]))
	    {
	        $tag = $user->lang[$tag_idx];
	    }
	    
		$template->assign_block_vars('custom_tags', array(
			'BBCODE_NAME'		=> "'[{$row['bbcode_tag']}]', '[/" . str_replace('=', '', $row['bbcode_tag']) . "]'",
			'BBCODE_ID'			=> $num_predefined_bbcodes + ($i * 2),
			'BBCODE_TAG'		=> $tag,
			'BBCODE_HELPLINE'	=> $helpline,
			'A_BBCODE_HELPLINE'	=> str_replace(array('&amp;', '&quot;', "'", '&lt;', '&gt;'), array('&', '"', "\'", '<', '>'), $helpline),
		));

		$i++;
	}
	
Since we want to access the language files, we also will have to import
the global ``$user`` variable in L811 which means replacing
    
    global $db, $template;
	
with 
    
    global $db, $template, $user;

within the same function.
	
When you now add a new custom BBCode -- let's call it "lala" here -- and you
want it's button to show something different depending on what language is
selected, simply add something like this to your language definitions in the
{$phpbb\_root\_path}/language/\*/posting.php:
    
    'CUSTOM_BBCODE_TAG_lala' => 'LALA button',
    'CUSTOM_BBCODE_HELP_lala' => 'Yet another useless button tooltip',
    
The ``CUSTOM_BBCODE_TAG_*`` entry defines the label of the button while the
``CUSTOM_BBCODE_HELP_*`` entry sets the tooltip. If you don't set them for
a language, the text you entered while creating the custom BBCode in the 
admin panel will be used.