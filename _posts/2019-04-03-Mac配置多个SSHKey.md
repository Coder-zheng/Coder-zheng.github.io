---

layout: post
title: "Mac配置多个SSHKey"
categories: 开发工具
tags:  Git
author: ant

---

* content
{:toc}
# 同一个Mac，配置多个SSH Key

## 缕清思路

网上关于这方面的文章已经挺多的了，但是每次要配置的时候，还是要靠搜索引擎找半天才搞定，索性自己记录一下。之前我遇到了下面几种情况：

* 自己github上项目，即要在自己的Mac上能提交代码，也要在公司的Mac上能提交代码。
* 公司的Mac上，要能往公司自己的git服务器上对应项目提交代码，也要能往自己github上的项目提交代码。
* 别人github上项目，在自己的Mac上能往上面提交代码。

解决上面的各种情况之前，先来说一下SSH Key是咋回事。公钥认证，是使用一对加密字符串，一个称为公钥（public key），任何人都可以看到其内容，用于加密；另一个称为密钥（private key），只有拥有者才能看到，用于解密。通过公钥认证可实现SSH免密码登陆，git的SSH方式就是通过公钥进行认证的。



也就是说，我们在自己的电脑上生成一对key，把其中的public key配置到git服务器的项目上，private key在自己的电脑上配置好，这时就是可以在自己的电脑上往git服务器上的那个项目提交代码了。哪怕你没有一个git帐号，但是能把这个public key配置到git服务器的项目上，比如你用两个鸡腿跪求别人把你的public key放上去，这时你照样能往上面提交代码。



现在来看一下上面的那几种情况需要怎么配置：

* 需要在自己的Mac上生成一对SSH Key，配置到自己的github项目上，在公司的Mac上也要生成一对SSH Key配置到自己的github项目上。
* 公司的Mac上，生成一对SSH key配置在公司自己的git服务器对应的项目上，再生成一个对SSH key配置在自己的github项目上。
* 在自己的Mac上生成一对SSH Key，配置在对方的github项目上。

## 开始配置

下面都拿github举例子了，毕竟大家都在用😁

#### 在github上的项目配置SSH Key中的public key

在github上创建一个repository后，在Settings中会找到如下项：

![image-20190403122103548](https://ws4.sinaimg.cn/large/006tKfTcly1g1pbqyk6kpj30h804sdgo.jpg)

这个Deploy keys就是配置可以往该项目上提交代码的public key，可以添加多个。

对于public key的配置就是这么简单，但是要注意，一个public key只能给一个repository用。

#### 在Mac上配置SSH Key中的private key

当往github的项目上提交代码时，github需要知道你电脑上有没有和那些Deploy keys中某个public key配对的private key。接下来就是配置怎样找到这个private key。

* 生成1个SSH Key:

```sh
$ ssh-keygen -t rsa -C "youremail@xxx.com"
```

按回车后：

```sh
Generating public/private rsa key pair.
Enter file in which to save the key (/Users/Shinancao/.ssh/id_rsa): id_rsa_TestSSH_github(取个名字)
Enter passphrase (empty for no passphrase): 
Enter same passphrase again: 
Your identification has been saved in 
id_rsa_TestSSH_github.
Your public key has been saved in 
id_rsa_TestSSH_github.pub.
```

最好每次生成时都给SSH Key取个名字，这样后面在管理时自己也一目了然。我这里的格式是id_rsa_项目名_git提供方，我生成的所有key都遵循这个规则命名。建议你也有你自己的一种命名方式，并且保持统一。如果不取名字，默认的是id_rsa，如果后面生成时不命名，会把这个覆盖掉。密码可以不设置，免得每次提交时还要输入一次，安全性自己衡量吧。第一次生成key时，会在~目录下创建一个.ssh目录。

```sh
$ cd ~/.ssh 
$ ls
```

* 把id_rsa_TestSSH_github.pub添加到github对应的项目的Deploy keys中。

![image-20190403122339844](https://ws4.sinaimg.cn/large/006tKfTcly1g1pbtlixvyj30h80bgtb5.jpg)

* ssh服务器默认是去找id_rsa，现在需要把这个key添加到ssh-agent中，这样ssh服务器才能认识id_rsa_TestSSH_github。

```sh
$ ssh-add -K ~/.ssh/id_rsa_TestSSH_github
```

这里为什么加上了一个-K参数呢？因为在Mac上，当系统重启后会“忘记”这个密钥，所以通过指定-K把SSH key导入到密钥链中。

查看添加结果：

```sh
$ ssh-add -l
```

* 创建本地的配置文件 ~/.ssh/config，编辑如下：

```sh
Host TestSSH.github.com
	HostName github.com
	User git
	PreferredAuthentications publickey
	IdentityFile ~/.ssh/id_rsa_TestSSH_github
Host YourProjectName.gitlab.com
	HostName gitlab.com
	User git
	PreferredAuthentications publickey
	IdentityFile ~/.ssh/id_rsa_YourProjectName_gitlab
```

Host的名字可以随意取，我这边按照的规则是项目名.git服务器来源，接下来会用到这个名字。测试是否配置正确：

```sh
$ ssh -T git@TestSSH.github.com
```

 (就是刚刚你给Host取的名字)

敲一下回车，如下出现下面的提示就连接成功了：

Hi shinancao/TestSSH! You've successfully authenticated, but GitHub does not provide shell access.
一定要注意哦，帐号名称/项目名称，如果这个key没有连接成功，它有可能提示的是别的key的。

修改github项目配置，使项目本身能关联到使用的key。
如果你在之前已经把项目clone到本地了，有两种解决方法：

(1) 打开项目目录/.git/config，将[remote “origin”]中的url中的github.com修改为TestSSH.github.com，就是你在第4步中给Host取的那个名字。如下：

remote "origin"]
	url = git@TestSSH.github.com:shinancao/TestSSH.git
	fetch = +refs/heads/*:refs/remotes/origin/*
(2) 也可以在提交时修改

$ git remote rm origin
$ git remote add origin git@TestSSH.github.com:shinancao/TestSSH.git
如果还没有clone到本地，则在clone时可以直接将github.com改为TestSSH.github.com，如下：

$ git clone git@TestSSH.github.com:shinancao/TestSSH.git
到此，就可以Happy Coding啦😄，可以push一次试试~

## github用户设置中的SSH Key

细心的小伙伴可能已经注意到了，在用户设置中也有一个SSH Keys的配置，这块添加的key是来设置一个电脑上默认使用的key的。每创建一个repository都弄一次Deploy Keys是挺麻烦的。



github默认找的是id_rsa这对密钥，所以此处要添加到github上的就是id_rsa.pub的内容。这对密钥一样存在于~/.ssh中，而且无需在config中设置。

先看一下id_rsa是否已经在ssh-agent中了：

$ ssh-add -l
如果不在要添加进去：

$ ssh-add -K ~/.ssh/id_rsa
测试是否能连接成功：

$ ssh -T git@github.com
敲一下回车，如果结果如下则成功了：

Hi shinancao! You've successfully authenticated, but GitHub does not provide shell access.
注意哦，这里只有用户名！没有跟着项目名了~ 配置完成后，就是可以轻松的创建repository，然后clone到本地，自由自在的往上面push代码啦~

## 配置邮箱和用户名

配置邮箱和用户名是用来干啥的呢？就是记录每一次commit的用户和与之关联的邮箱。可以在电脑上配置一个全局的user.name和user.email，也可以针对不同的repository配置不同的user.name和user.email。



配置全局的user.name和user.email：

$ git config --global user.name "your name"
$ git config --global user.email "your email"
查看结果：

$ git config --global user.name
$ git config --global user.email
位置在~/.gitconfig文件中。

在做这块的测试时，我发现了一个很有意思的事情，对于github，即使我随意设置了一个全局的name，但最后提交完，显示的还是我github的用户名。当取消了global的设置，只针对某个repository设置，则github上可以显示我设置的了。

如果同时设置了全局的，和针对某个repository的，则优先使用全局的。那要单独给每个repository设置，要先取消全局的设置。

$ git config --global --unset user.name
$ git config --global --unset user.email
然后进入到项目目录下设置：

$ git config user.email "your name"
$ git config user.email "your email"
查看结果：

$ git config user.name
$ git config user.email
位置在项目目录/.git/config文件中。

## HTTPS的方式提交代码

通过https的方式clone url，然后再以https的方式提交代码，我觉得就是正常的使用帐号和密码去操作。这种方式需要你知道项目拥有者的帐号和密码，而且在每次commit时都要输入用户名和密码。显然不方便啦，尤其是你要参与到别人的项目中去开发，人家总不能把帐号名和密码都给你吧😂。