<?php
$host = "192.168.1.176";
$port = 6808;

$client = new ClientExt($host, $port);
$client->connect();

$text = "南京大学城书店，长春市长春药店，研究生命起源";
$params = msgpack_pack(array($text));

$client->dowork("PartWordsM1", $params, function($ret) {
	var_dump($ret);
});

$client->dowork("PartWordsM2", $params, function($ret) {
	var_dump($ret);
});

$client->dowork("PartWordsM3", $params, function($ret) {
	var_dump($ret);
});

$client->close();