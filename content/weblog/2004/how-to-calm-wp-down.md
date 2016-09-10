---
date: '2004-11-23T12:00:00-00:00'
language: en
tags:
- wordpress
title: How to calm WP down ;-)
---


WordPress is really a very versatile and extensible system, but who wants to integrate into an already existing application will inevitably get to a point, were the echo'ing template functions start to get on your nerves. This method has some advantages for example...
* is it easier to integrate these functions into a HTML site (ok, one char less work if you use the short echo method of PHP (`=`)

On the other hand this method also has the disadvantage that you can't simply use the result of these functions for other purpose for example for passing it to the functions of your already existing application. Also it's harder to control the actual output.

-------------------------------



You have basically 3 options to solve this problem and to get the result of these functions as a return value instead of seeing it appear on your site right after calling the function:
1. Rewrite the original WP template functions
2. Write your own template functions 
3. Write a wrapper around the WP template functions using the output buffering functions of PHP

(1) and (2) are not really useful on the long run simply because you'd have to check or rebuild them everytime a new update of WP appears in the net. So I'll explain (3) here in detail and post a short but still quite powerful wrapper function for the WP template function that should make your life much easier ;-) ... and no, it won't help you with your paypal problems ;-)

My wrapper function is quite simple but should be able to help in nearly every situation were some output redirection of a function is necessary.

<pre class="code">
function o2r()
{
   global $_lastreturn;
   $_lastreturn = NULL;
   if(1 > func_num_args())
   {
      trigger_error("o2r called without any argument",E_USER_ERROR);
      return "";
   }
   else
   {
      /* All the errors should basically already be handled
       * by php itself */
      $args = func_get_args();
      $func = array_shift($args);
      ob_start();
      $_lastreturn = call_user_func_array($func,$args);
      $ret = ob_get_contents();
      ob_end_clean();
      return $ret;
   }
}</pre>

Let's make a small example here now. The function we would like to use is this one:
<pre class="code">function echoing()
{
   echo "Do you hear me?";
   return 1;
}</pre>

It echos the string "Do you hear me?" and returns the integer 1. Well, but we don't want the function to echo "Do you hear me?" but instead get this string as a return value. So we can call this function with our o2r wrapper function

<pre class="code">$output = o2r('echoing');</pre>

Now we have "Do you hear me?" in the $output instead of on the screen and the 1 that was returned by the echoing function is now stored in $GLOBALS['_lastreturn'].

Can you also use o2r if your function needs parameters? Sure you can :-) Simply pass them after the function name:

<pre class="code">o2r('echoing2',$arg1,$arg2);</pre>

If you have any suggestions please let me know :-) I hope this function will help some people out there :-)