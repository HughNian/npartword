<div align="center">
    <a href="http://www.niansong.top"><img src="https://raw.githubusercontent.com/HughNian/npartword/master/logo/npartword_logo1.png" alt="npw logo" width="160"></a>
</div>


## npw介绍
npw：npartword，golang实现中文分词系统，主体分词逻辑有两个部分。   

- 1.前缀树查找字典，通过disctance或mmseg算法过滤分词    

- 2.维特比算法解隐马尔可夫模型，对词进行隐状态标注分词   

- 3.加入情感词典，文本情感分类算法，对文本进行情感评分，如积极，消极，否定等，一般分值越高文本的积极性越高   

分词服务系统的实现，通过nmid的worker方式实现。       

- 1.分词系统服务端，需要实现nmid的worker，服务的实现十分简单，无需考虑通信问题，这些nmid解决。   

- 2.分词系统服务调用，只要通过nmid的client调用即可，任何nmid的client都可以随时跨服务器的使用分词系统。    

## k8s使用
在k8s目录中有k8s部署文件，在k8s集群中执行`sudo kubectl apply -f deployment.yaml`，使用nmid镜像#56，nsearch镜像#5，支持链路追踪。   
在k8s部署skywalking，在k8s集群中执行 `sudo kubectl apply -f skywalking-oap.yaml`，`sudo kubectl apply -f skywalking-ui.yaml`
<div align="center">
    <img src="https://raw.githubusercontent.com/HughNian/npartword/master/trace.png" alt="trace">
</div>  


```go
php调用示例

$host = "xxx.xxx.x.xx";
$port = xxx;
$client = new ClientExt($host, $port);
$client->connect();

//普通分词
//fname - PartWordsM1
$text = "南京大学城书店，长春市长春药店，研究生命起源";
$params = msgpack_pack(array($text, "1"));
$client->dowork("PartWordsM1", $params, function($ret){
        var_dump($ret);
});

array(3) {
  [0]=>
  int(0)
  [1]=>
  string(2) "ok"
  [2]=>
  string(77) "南京大学|城|书店|，|长春市|长春|药店|，|研究|生命|起源|"
}

//mmseg分词
//fname - PartWordsM2
$params = msgpack_pack(array($text, "1"));
$client->dowork("PartWordsM2", $params, function($ret){
        var_dump($ret);
});

array(3) {
  [0]=>
  int(0)
  [1]=>
  string(2) "ok"
  [2]=>
  string(77) "南京|大学城|书店|，|长春市|长春|药店|，|研究|生命|起源|"
}

//隐马尔可夫模型
//fname - PartWordsM3
$params = msgpack_pack(array($text, "1"));
$client->dowork("PartWordsM3", $params, function($ret) {
        var_dump($ret);
});

array(3) {
  [0]=>
  int(0)
  [1]=>
  string(2) "ok"
  [2]=>
  string(75) "南京大学|城书店|，|长春市|长春药店|，|研究|生命|起源|"
}

//获取文本情感值
$text2 = "世事无常，你是人间琳琅，众人平庸，你是人间星光，万事浮沉，你是人间归途。";
$params = msgpack_pack(array($text, "3"));
$client->dowork("PartWordsM1", $params, function($ret){
        var_dump($ret);
});

array(3) {
  [0]=>
  int(0)
  [1]=>
  string(2) "ok"
  [2]=>
  string(213) "世事/x| 无常/ne| ，/x| 你/x| 是/po| 人间/x| 琳琅/x| ，/x| 众人/x| 平庸/x| ，/x| 你/x| 是/po| 人间/x| 星光/x| ，/x| 万事/x| 浮沉/x| ，/x| 你/x| 是/po| 人间/x| 归途/x| 。/x| =2.00"
  //获取情感值分词会带上情感词性，最后会有情感值分数
}

$text2 = "今天我想你了 但是这不影响我打游戏 不影响我睡觉 只是悄咪咪抹了抹眼泪 感叹了一下最后还是没能牵到你的手";
$params = msgpack_pack(array($text, "3"));
$client->dowork("PartWordsM1", $params, function($ret){
        var_dump($ret);
});

array(3) {
  [0]=>
  int(0)
  [1]=>
  string(2) "ok"
  [2]=>
  string(319) "今天/x| 我/x| 想/po| 你/x| 了/x|  /x| 但是/x| 这/x| 不/ga| 影响/x| 我/x| 打游戏/x|  /x| 不/ga| 影响/x| 我/x| 睡觉/x|  /x| 只是/x| 悄/po| 咪咪/x| 抹/x| 了/x| 抹/x| 眼泪/x|  /x| 感叹/x| 了/x| 一下/x| 最后/x| 还是/x| 没/ga| 能/po| 牵/x| 到/po| 你/x| 的/x| 手/x| 。/x| =0.50"
}
//同样是表达爱意，显然这句的情感不如上一句，所以他的分值低一些
```

## 交流博客
http://www.niansong.top