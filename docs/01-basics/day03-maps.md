# Day 3: Map实战应用

## 学习目标
- 深入理解Map的底层原理
- 掌握Map在实际项目中的应用场景
- 学会处理并发安全问题
- 实现简单的缓存系统

## 核心概念

### 1. Map基础回顾
- Map是Go中的哈希表实现
- key必须是可比较类型
- value可以是任意类型
- 零值是nil，使用前必须初始化

### 2. Map的实际应用场景
- **配置管理**：存储应用配置
- **缓存系统**：临时存储计算结果
- **计数器**：统计词频、访问次数
- **索引构建**：快速查找数据

## 今日练习

### 练习1：实现一个简单的内存缓存
创建 `exercises/01-basics/day03-maps/cache.go`

要求：
1. 实现Set(key, value, ttl)方法
2. 实现Get(key)方法
3. 支持过期时间（TTL）
4. 实现Delete(key)方法
5. 实现Clear()清空缓存

提示：
- 使用map[string]interface{}存储数据
- 需要额外存储过期时间
- 考虑使用time.AfterFunc或定期清理

### 练习2：实现配置管理器
创建 `exercises/01-basics/day03-maps/config.go`

要求：
1. 支持从map加载配置
2. 实现Get(key)获取配置值
3. 实现GetString, GetInt, GetBool等类型安全方法
4. 支持嵌套配置（如"database.host"）

### 练习3：实现并发安全的计数器
创建 `exercises/01-basics/day03-maps/counter.go`

要求：
1. 实现Increment(key)增加计数
2. 实现Get(key)获取计数
3. 必须是并发安全的
4. 实现GetAll()返回所有计数

## 实战任务：构建简单的Session管理器
创建 `exercises/01-basics/day03-maps/session.go`

结合今天所学，实现一个Web应用中常用的Session管理器：
1. 创建Session（返回sessionID）
2. 存储和获取Session数据
3. Session过期处理
4. 并发安全

## 注意事项
1. Map不是并发安全的，需要加锁或使用sync.Map
2. 遍历Map的顺序是随机的
3. 删除不存在的key不会panic
4. 注意内存泄漏（特别是缓存场景）

## 扩展阅读
- sync.Map的使用场景
- Map的底层实现原理
- 如何优化Map的性能

## 明日预告
Day 4: 接口设计实战 - 设计一个插件系统