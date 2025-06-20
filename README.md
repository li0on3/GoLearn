# GoLearn - Go后端开发学习项目

这是一个系统化的Go语言后端开发学习项目，注重实践和规范化。

## 项目结构
- `docs/` - 学习文档和笔记
- `exercises/` - 各阶段练习代码
- `projects/` - 实战项目
- `CLAUDE.md` - AI助手交接文档

## 学习目标
培养扎实的Go语言基础，掌握后端开发核心技能，养成规范的编程习惯。

## 如何使用
1. 查看 `docs/00-overview.md` 了解完整学习计划
2. 按照文档指引逐步学习
3. 所有代码需要自己动手编写
4. 定期查看 `docs/progress.md` 跟踪进度

## Git工作流
```bash
# 每次学习前
git checkout -b feature/day-xx-topic

# 完成练习后
git add .
git commit -m "feat: 完成xxx练习"
git checkout main
git merge feature/day-xx-topic
```

## 开始学习
请从 `docs/01-basics/day01-setup.md` 开始您的Go语言学习之旅！