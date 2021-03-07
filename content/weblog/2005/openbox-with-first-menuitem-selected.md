---
date: '2005-05-21T12:00:00-00:00'
language: en
tags:
- linux
title: OpenBox WITH first menuitem selected
---


If you're using the CVS version of OpenBox or the latest release candidate for 3.3 you've perhaps noticed that the first menu in the root-menu is no longer auto-selected when opening this menu. I think this new behaviour was introduced for people having big submenus right on top in their menus. This could result in the menu jumping around on the screen to find some way to display all the content of the submenu. 

-------------------------------



But if you're like me , you'll perhaps have a terminal emulator instance as the first element in the root menu which now takes one keystroke more to be activated. After 2 days trying to get used to this I grew  quite desperated hitting enter half of the time in vain... So I tried to find the actual change in the source code and found it in openbox/menu.c. This small patch should undo this change:



<pre class="code">

diff -NurBabdH openbox/openbox/menu.c openbox-r1/openbox/menu.c

--- openbox/openbox/menu.c      2005-05-21 00:25:19.000000000 +0200

+++ openbox-r1/openbox/menu.c   2005-05-21 00:21:57.000000000 +0200

@@ -307,6 +307,11 @@

                         x - ob_rr_theme->bwidth, y - ob_rr_theme->bwidth);

     if (!menu_frame_show(frame, NULL))

         menu_frame_free(frame);

+    else {

+           if (frame->entries){

+                   menu_frame_select_next(frame);

+           }

+    }

 }

 

 static ObMenuEntry* menu_entry_new(ObMenu *menu, ObMenuEntryType type, gint id)

</pre>



For a __real__ solution I'd suggest that you vote for <a href="https://bugzilla.icculus.org/show_bug.cgi?id=2228">this bug</a> :-)
