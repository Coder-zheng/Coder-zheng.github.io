---
layout: page
title: About
permalink: /about/
icon: heart
type: page
---

* content
{:toc}
## 联系方式

- 手机：        15510369979
- Email：       cqmyg14ds2@gmail.com
- 微信：        15510369979

## 个人信息

- 基本信息： 郑俊领/男/1990.04.28
- 工作年限： 4年半（2年 Android ，2年半 Golang 区块链）
- 学历：        本科/英语专业
- 技术博客： https://coder-zheng.github.io/
- Github：    https://github.com/Coder-zheng
- 期望职位： 区块链开发工程师
- 期望城市： 北京
- 自我评价： 爱学习，熟练掌握区块链，Golang 开发， Android 端开发，熟悉前端开发，沉稳冷静，认真负责，办事靠谱，善于沟通，积极主动，吃苦耐劳，对技术充满激情。

## 工作经历

### 久康国际健康股份有限公司（2015年7月~2018年11月）

#### 一、9K 药品溯源

**项目介绍**：

基于 Fabric 技术快速搭建Hyperledger fabric区块链网络，对联盟内企业库存管理系统进行整合，实现药品的溯源，将药品生产商、药品采购商、药品销售商、监管部门等组织加入联盟链之中,从生产到客户购买过程的透明,公平化。本项目使用 Beego 框架实现B/S后台的快速搭建，基于 Fabric 实现核心溯源功能。

**使用的技术**：

1. 编写 Yaml 配置文件,完成对区块链网络的快速搭建，使用 Cryptogen 、Configtxgen 工具生成内部证书，创始块，通道等
2. 根据业务对 Fabric 的 Chaincode 链码的编写，通过 Docker-compose 对容器进行编排启动,通过对 Ca 服务器交互来实现新用户加入到当前区块链网络 
3. 对 Fabric-go-sdk 进行 Fabric 功能封装，提供所需接口
4. 对配置文件进行修改并进行 Kafka 集群的搭建
5. 根据需求文档进行 Beego 后台开发:组织管理，角色管理，权限管理，数据字典管理，物品入库管理等功能

#### 二、9K 患者爱心筹项目

**项目介绍**：

一款基于以太坊开发的患者众筹项目，解决了筹款人与捐款人之间的信任问题，公开透明了每一笔捐款的去向，同时利用 IPFS 分布式文件存储系统，实现了患者病历的永久保存，各医院之间的同步问题。

**使用的技术**：

1. 使用 Remix 编写 Solidity 智能合约
2. 使用 Ganache 配合 Metamask 和 Mocha 进行测试
3. 使用 Solc 编译合约，Web3.js 部署合约
4. 使用 Truffle 框架进行智能合约快速编译部署
5. 使用 IPFS 实现图片资源的永久存储
6. 使用 React，Js，Html 快速搭建前台页面

#### 三、9K 公链开发

**项目介绍**：

一款简化版的加密货币，实现了工作量证明，区块持久化，区块链迭代器，命令行接口， Coinbase 交易，普通转账交易， UTXO 模型，加锁解锁脚本，公钥生成货币地址，Wallets 保存多个密钥对等功能。

**使用的技术**：

Bitcoin 原理、Golang、密码学、BoltDB

#### 四、9K 医疗信息管理系统

**项目介绍**：

医院信息管理，医生信息管理，患者信息管理。其中医生模块包含，账户注册，医生排版查询，医生登录，修改个人信息，修改登录密码，患者信息管理，添加患者就诊信息，患者就诊信息查询，添加用药信息，查看用药信息，治疗检查管理。其中患者模块包含，账户注册，医生信息查询，添加病历，修改病历，在线挂号，在线问诊，在线评价。

**使用的技术**：

1. GoMicro 搭建服务架构、Consul 作为服务发现，其中采用 RESTful 协议与 gRPC/Protobuf 作为微服务间通信机制
2. Beego快速搭建后台服务，WebSocket实现简单的网页聊天
3. 使用 Mysql 做数据持久化，以及 Redis 提升项目性能
4. 使用 Nginx 做负载均衡，FastDfs 实现文件的分布式存储

#### 五、9K 医生端，患者端 App（前后端都参与开发） 

**项目介绍**：

一套线上移动问诊 App ，分医生端患者端。响应国家分级诊疗政策，缓解一线城市医疗资源压力，为解决看病难问题。患者端，可以登录注册，实名认证，修改密码，添加就诊人/家庭成员，上传病历/修改病历，获取医生详情/医院详情，微信支付宝充值，创建订单，视频聊天，图文聊天。

**使用的技术**：

- Server 端：Golang 、微服务、Grpc 、Docker 、Nginx 、Beego 、Mysql 、Redis 、Nginx 、FastDfs 、RESTful 、 声网
- 安卓 App 端：Java 、Mvp 、自定义 View 、事件传递机制、多线程下载、Okgo 、Bugly 、Eventbus 通讯、Realm 数据库、声网视频聊天、环信即时通讯、友盟数据统计

### 北京竞技时代科技有限公司（2014年7月~2015年7月）

#### 六、WVA 赛事 App

**项目介绍**：

一款 WVA 赛事体系下的 App ，类似掌上英雄联盟，实现了 VR 设备扫码登录，赛事举办查询报名一条龙服务，选手或消费者，在线选择线下体验店，并领取各种优惠券，在线微信，支付宝支付，还可以在线聊天视频

**使用的技术**：

Java 、Mvp 、Http 、Tcp 、Socket 、自定义 View 、事件传递机制、多线程下载、Realm 数据库、声网视频聊天、友盟数据统计

## 技能清单

1. 精通 Golang 开发
2. 熟练掌握安卓开发，测试，上线，即时通讯，视频，支付等
3. 熟练掌握 Golang 实现 Bitcoin 公链开发
4. 熟练掌握 Hyperledger Fabric 相关机制及网络的部署以及链码的编写
5. 熟练掌握 对称、非对称加密原理，POW，POS，ECC 等基本使用
6. 熟练理解 Raft，PoW，PoS，Paxos，DPos，2PC，3PC 等共识算法
7. 熟练掌握 Goroutine channel，网络协议，网络编程，高并发服务器
8. 熟练掌握 Beego 框架，MVC 设计模式，ORM 对象操作数据库数据
9. 熟练掌握关系型数据库 Mysql，非关系型 Redis 的基本操作
10. 熟练掌握 Web3.js，Truffle，Ipfs 等使用
11. 熟练掌握 Solidity 智能合约开发以及部署测试
12. 熟练掌握 Docker 的基本使用和 Dockercompose 进行项目部署
13. 熟练掌握 Linux，Mac，Windows，Git，IPFS 基本操作
14. 了解 Go-micro ，gRPC框架进行 Go 语言微服务编写
15. 了解 Nginx，Kafka，Zookeeper 等环境配置
16. 了解 React ，Css ，Html ，原生 Js前端开发
17. 熟练掌握 Solidity ，熟悉 Java ，熟悉原生 Js ，Nodejs ，Python（写过一些爬虫），Shell（可写简单脚本），Kotlin（用 Kotlin 对 Java 项目做过重写），C ，C++（能够读懂）