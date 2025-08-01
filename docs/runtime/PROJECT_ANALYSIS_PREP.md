# 项目分析所需信息 (Information Required for Project Analysis)

为了对 `origadmin/runtime` 项目进行全面且准确的分析，请您尽可能详细地填写以下信息。项目定位：**一个基于 go-kratos
的、面向分布式架构的Go后端快速开发框架**。

---

### 1. 项目核心目标 (Core Project Objective)

- **一句话描述项目**：一个基于 go-kratos 的、面向分布式架构的Go后端快速开发框架，旨在提升微服务开发效率和规范性，并使开发者无需直接操作 `go-kratos` 的底层 API 即可完成项目开发。
- **主要用户/场景**：
    - **面向开发者/团队**：中小型团队、需要快速构建微服务应用的开发者、对 `go-kratos` 有一定了解并希望在此基础上获得更统一开发体验的团队。
    - **解决问题**：降低微服务开发门槛、提供统一的开发规范和最佳实践、加速项目启动、减少重复性工作、简化 `go-kratos` 的使用复杂性。
- **外部关联**：
    - `Toolkits`：提供**独立于 `runtime` 核心逻辑的、可复用的通用工具函数和库**。`Runtime` 可以依赖 `Toolkits`，但 `Toolkits` 不应依赖 `Runtime`。
    - `Runtime`：作为**核心框架**，提供统一的接口、组件和运行时环境，供上层应用（通过项目生成工具）使用。`Runtime` 可以依赖 `Toolkits` 和 `Contrib`。
    - `Contrib`：提供**对 `Runtime` 接口的第三方库实现或扩展**。例如，`Contrib/database` 提供 `runtime/interfaces/storage.go` 接口的 MySQL/PostgreSQL 实现。`Contrib` 依赖 `Runtime` 的接口定义，并通过依赖注入的方式在运行时组装，`Runtime` 不直接依赖 `Contrib` 以避免循环引用。

---

**### 2. 开发与维护状态 (Development & Maintenance Status)

- **当前阶段**：(请选择)
    - [x] 早期原型/概念验证 (Early Prototype/PoC)
    - [x] 积极开发中 (In Active Development)
    - [ ] 已上线/稳定维护 (Production/Stable Maintenance)
    - [ ] 遗留系统/仅修复关键Bug (Legacy System/Critical Bugs Only)
- **团队痛点**：当前开发或维护此框架时，遇到的最大挑战是什么？
    - 这是一个全新的构思, 因为是从零开始，没有历史参考，因此需要考虑很多东西。
    - 还有需要很多关联工具需要开发, 比如一键生成项目结构，一键生成代码，一键生成配置文件等等。
    - 还有通用接口的设计需要考虑到兼容性, 所以需要考虑很多东西。
    -

---

### 3. 架构与设计 (Architecture & Design)

- **架构风格**：本项目 `origadmin/runtime` 定位为一个**基于 go-kratos 的、面向分布式架构的Go后端快速开发框架**。其架构风格旨在提供**统一的、可扩展的公共接口和组件**，以支持项目生成和快速开发。我们通过以下方式实现：
    - **公共 API 优先：** 框架的核心功能和扩展点以 Go 包的形式直接暴露在 `runtime` 模块的顶层目录，而非隐藏在 `internal` 中，以便外部模块（如生成的项目）可以直接导入和使用。
    - **接口驱动设计：** 关键抽象（如 `Service`, `Registry`, `Middleware` 等）以 Go 接口的形式定义在 `runtime/interfaces` 目录下，实现依赖倒置，允许底层实现灵活替换。
    - **适配与扩展：** 在 `go-kratos` 的基础上，我们对现有抽象进行适配和封装，并针对其不支持的功能进行扩展，以满足更广泛的业务需求和统一开发体验。
    - **目录结构：**
        - `api/`: 定义外部通信协议（如 Protobuf）。
        - `bootstrap/`: 框架的标准化启动和依赖组装逻辑。
        - `config/`: 框架的统一配置系统。
        - `context/`: 框架的自定义上下文或上下文工具。
        - `interfaces/`: **框架对外暴露的所有统一 Go 接口定义。**
        - `log/`: 框架的统一日志接口和实现。
        - `mail/`: 框架的统一邮件发送接口和实现。
        - `middleware/`: 框架的通用中间件组件。
        - `registry/`: `IRegistry` 接口的具体实现。
        - `service/`: `IService` 接口的具体实现。
        - `storage/`: `IStorage` 接口的具体实现。
        - `third_party/`: 第三方 proto 文件或其他公共依赖。
        - `internal/`: **仅供 `runtime` 模块内部使用的私有辅助代码，不应被外部导入。**

- **核心抽象**：框架的核心抽象是 `runtime` 模块对外提供的**统一接口和可插拔组件**。这些抽象与 `go-kratos` 的原生抽象关系如下：
    - **`Service` (位于 `runtime/service/`，实现 `runtime/interfaces/service.go`):** 继承/适配 `go-kratos` 的 `Service` 抽象，并添加 `runtime` 框架特有的扩展功能和统一接口。
    - **`Registry` (位于 `runtime/registry/`，实现 `runtime/interfaces/registry.go`):** 继承/适配 `go-kratos` 的 `Registry` 抽象，并添加 `runtime` 框架特有的扩展功能和统一接口。
    - **`Middleware` (位于 `runtime/middleware/`):** 继承/适配 `go-kratos` 的 `Middleware` 抽象，并提供 `runtime` 框架统一的中间件实现。
    - **`Config` (位于 `runtime/config/`):** 封装 `go-kratos` 的配置机制，并提供 `runtime` 框架统一的配置加载和管理接口。
    - **`Context` (位于 `runtime/context/`):** 定义 `runtime` 框架的运行时上下文，可能包含对 `go-kratos` 上下文的扩展。
    - **`Log` (位于 `runtime/log/`):** 继承/适配 `go-kratos` 的 `Log` 抽象，并提供 `runtime` 框架统一的日志接口和实现。
    - **`Bootstrap` (位于 `runtime/bootstrap/`):** `runtime` 框架的标准化启动组件，负责组装所有依赖，可能基于 `go-kratos` 的启动流程进行扩展。
    - **`Interface` (位于 `runtime/interfaces/`):** 这是 `runtime` 框架的核心，定义了所有对外暴露的统一 Go 接口，供使用者依赖。
    - **`Storage` (位于 `runtime/storage/`，实现 `runtime/interfaces/storage.go`):** `runtime` 框架提供的统一数据访问接口和实现，用于封装底层数据库或存储操作。
    - **`Mail` (位于 `runtime/mail/`，实现 `runtime/interfaces/mail.go`):** `runtime` 框架提供的统一邮件发送接口和实现.
- **计划集成的组件**：框架计划或已经集成了哪些核心的第三方服务/组件作为可选项？
    - 考虑到`Runtime`本身是一个通用的运行时框架，因此，我们是否需要添加额外的组件呢？
    - 
---

### 4. 关键非功能性需求 (Key Non-Functional Requirements)

- **可扩展性**：框架的关键模块（如：配置、注册、中间件）设计了清晰的扩展点（plugin/interface）。通过提供清晰的接口和示例，旨在降低开发者自定义扩展的难度。
    - 针对`Service`组件, 后续需要支持当前的主流框架库, 如`Gin`、`Echo`、`Fiber`等。
- **性能要求**：框架本身引入的性能开销（overhead）是否有评估或目标？
    - 在可控范围内, 尽可能使用性能较高的实现方案.但是稳定性、可扩展性、可维护性也是重要的考虑因素。
- **可观测性**：项目将利用并增强 `go-kratos` 的可观测性能力。确保项目在可观测性方面有优秀的实践，例如，提供统一的日志格式、集成 OpenTelemetry 进行链路追踪和指标收集，并提供默认的仪表盘配置建议。
---

### 5. 未来规划 (Future Roadmap)

- **近期功能**：未来3-6个月内，计划添加哪些主要功能或进行哪些重大重构？
    - 完成核心接口（Service, Registry, Storage, Config, Log）的初步实现和 `go-kratos` 适配。
    - 开发项目生成工具的 MVP 版本。
    - 完成 Gin/Fiber 框架的 `Service` 组件适配。
- **长期愿景**：希望这个框架最终发展成什么样子？在技术社区或公司内部扮演什么样的角色？
    - 目标是用户可以通过简单的配置, 快速搭建出一套完整的基于微服务的应用系统。
    - 成为 Go 微服务开发的行业标准之一，拥有活跃的社区和丰富的生态系统，支持多种基础设施和部署环境。

---