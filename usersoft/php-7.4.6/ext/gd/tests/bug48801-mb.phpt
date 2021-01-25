--TEST--
Bug #48801 (Problem with imagettfbbox)
--SKIPIF--
<?php
	if(!extension_loaded('gd')){ die('skip gd extension not available'); }
	if(!function_exists('imageftbbox')) die('skip imageftbbox() not available');
	if (substr(PHP_OS, 0, 3) == 'WIN') die('skip UTF-8 font file names not yet supported on Windows');
?>
--FILE--
<?php
$cwd = __DIR__;
$font = "$cwd/Tuffy私はガラスを食べられます.ttf";
$bbox = imageftbbox(50, 0, $font, "image");
echo '(' . $bbox[0] . ', ' . $bbox[1] . ")\n";
echo '(' . $bbox[2] . ', ' . $bbox[3] . ")\n";
echo '(' . $bbox[4] . ', ' . $bbox[5] . ")\n";
echo '(' . $bbox[6] . ', ' . $bbox[7] . ")\n";
?>
--EXPECTREGEX--
\(4, 15\)
\(16[0-1], 15\)
\(16[0-1], -4[7-8]\)
\(4, -4[7-8]\)
