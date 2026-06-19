# 08 — Java / Spring Boot

> Focus on JVM internals, Spring Boot patterns, and concurrency. Kotlin where relevant.

---

## JVM Internals

**Key Terms:**
- `JVM (Java Virtual Machine)` — executes Java bytecode; platform-independent
- `bytecode` — compiled .class file; not machine code; interpreted/JIT-compiled by JVM
- `JIT (Just-In-Time) compiler` — compiles hot bytecode to native machine code at runtime
- `HotSpot` — Oracle's JVM; most common; identifies "hot" code paths for JIT compilation
- `class loader` — loads .class files; Bootstrap → Extension → Application ClassLoader chain
- `method area (Metaspace)` — stores class metadata, static variables, method bytecode (Java 8+ no longer uses PermGen)
- `heap` — objects and arrays; managed by GC; divided into young/old generations
- `stack` — per-thread; holds stack frames (local variables, operand stack, frame data)
- `stack frame` — one per method call; contains local variables, return address
- `PC register` — per-thread program counter; holds current instruction address
- `native method stack` — for native (C/C++) method calls (JNI)
- `StackOverflowError` — recursion too deep; stack exhausted
- `OutOfMemoryError` — heap exhausted; can be heap OOM or Metaspace OOM
- `JVM flags` — `-Xmx` (max heap), `-Xms` (initial heap), `-XX:+UseG1GC` (GC algorithm)

**Must understand:**
- [ ] JVM memory areas — heap vs stack vs metaspace (what goes where)
- [ ] Why Java is "write once, run anywhere" (bytecode + JVM abstraction)
- [ ] JIT compilation — why JVM starts slow but gets faster (warmup)
- [ ] -Xmx vs -Xms — why starting small is sometimes wrong in production

---

## Java Garbage Collection

**Key Terms:**
- `generational hypothesis` — most objects die young; GC optimizes for this
- `young generation` — new objects; collected frequently; fast (minor GC)
- `eden space` — new objects created here
- `survivor spaces (S0, S1)` — objects that survive eden GC
- `old generation (tenured)` — objects that survive multiple minor GCs; collected rarely
- `minor GC` — collects young generation; fast (milliseconds)
- `major GC / full GC` — collects old generation; slow (pause entire application)
- `stop-the-world pause` — all application threads pause during GC; source of latency spikes
- `Serial GC` — single-threaded GC; `-XX:+UseSerialGC`; only for small apps
- `Parallel GC` — multi-threaded throughput-focused; `-XX:+UseParallelGC`; Java 8 default
- `G1 GC (Garbage First)` — region-based; predictable pause times; Java 9+ default; `-XX:+UseG1GC`
- `ZGC` — ultra-low pause (<10ms); works with huge heaps; Java 15+; `-XX:+UseZGC`
- `Shenandoah` — concurrent compaction; RedHat; very low pause; `-XX:+UseShenandoahGC`
- `GC roots` — starting points for marking: local variables, static fields, JNI references
- `mark phase` — traverse object graph from GC roots; mark live objects
- `sweep phase` — free unmarked objects
- `compaction` — move live objects together to eliminate fragmentation
- `memory leak in Java` — holding unnecessary references (static collections, listeners not unregistered)

**Must understand:**
- [ ] Why generational GC is efficient (most objects are short-lived)
- [ ] Minor GC vs major GC — frequency and pause time difference
- [ ] G1 GC — why it's preferred (predictable pauses, handles large heaps)
- [ ] How memory leaks happen in GC languages (strong references prevent collection)

---

## Java Concurrency

**Key Terms:**
- `Thread` — class; `new Thread(runnable).start()`
- `Runnable` — functional interface; `run()` method; no return, no checked exceptions
- `Callable<T>` — functional interface; `call()` method; returns T, can throw checked exceptions
- `ExecutorService` — thread pool; submit tasks; shutdown gracefully
- `Executors.newFixedThreadPool(n)` — pool of n threads
- `Executors.newCachedThreadPool()` — grows/shrinks dynamically; can exhaust resources
- `Executors.newSingleThreadExecutor()` — one thread; tasks queued sequentially
- `Future<T>` — represents pending result; `get()` blocks; `cancel()`, `isDone()`
- `CompletableFuture<T>` — async composition; `thenApply`, `thenCompose`, `allOf`, `anyOf`
- `synchronized` — keyword; acquire monitor lock on object; one thread at a time
- `ReentrantLock` — explicit lock; `lock()`, `unlock()`; supports tryLock, fairness
- `ReentrantReadWriteLock` — separate read and write locks
- `volatile` — variable visible to all threads immediately; no caching in CPU registers; not atomic for compound ops
- `AtomicInteger, AtomicLong` — lock-free atomic operations (CAS-based)
- `ThreadLocal<T>` — per-thread variable storage; useful for request context
- `CountDownLatch` — count down to zero; await until zero (one-time barrier)
- `CyclicBarrier` — wait for N threads to reach barrier; reusable
- `Semaphore` — permits-based; `acquire()` / `release()`; limit concurrency
- `BlockingQueue` — thread-safe queue with blocking put/take
- `ConcurrentHashMap` — thread-safe map; segment locking; better than synchronized HashMap
- `virtual threads (Java 21)` — lightweight like goroutines; `Thread.ofVirtual().start(runnable)`

**Must understand:**
- [ ] synchronized vs ReentrantLock — when to use each (synchronized is simpler; lock enables tryLock/fairness)
- [ ] CompletableFuture composition — thenApply (sync transform) vs thenCompose (async)
- [ ] Why ConcurrentHashMap is better than `Collections.synchronizedMap()`
- [ ] volatile vs synchronized — volatile for visibility only; synchronized for atomicity + visibility
- [ ] Java 21 virtual threads — how they compare to goroutines

---

## Spring Boot Basics

**Key Terms:**
- `Spring Framework` — IoC container + AOP + data access + web layer
- `Spring Boot` — opinionated Spring; auto-configuration; embedded server; no XML
- `auto-configuration` — Spring Boot automatically configures beans based on classpath and properties
- `starter dependencies` — `spring-boot-starter-web` pulls in Tomcat, MVC, Jackson automatically
- `application context` — Spring's IoC container; holds all beans; `ApplicationContext`
- `@SpringBootApplication` — combines @Configuration + @EnableAutoConfiguration + @ComponentScan
- `@ComponentScan` — scans packages for @Component, @Service, @Repository, @Controller
- `embedded server` — Tomcat (default), Jetty, Undertow; no separate server installation needed
- `application.properties / application.yml` — configuration; override any Spring defaults
- `profiles` — `spring.profiles.active=prod`; different configs per environment
- `actuator` — production-ready endpoints: /health, /metrics, /env, /beans; add `spring-boot-starter-actuator`
- `banner` — ASCII art on startup (cosmetic; can be disabled)
- `CommandLineRunner / ApplicationRunner` — run code after context starts

**Must understand:**
- [ ] How auto-configuration works — conditional beans (@ConditionalOnClass, @ConditionalOnProperty)
- [ ] Why Spring Boot is "opinionated" — sensible defaults, less boilerplate
- [ ] How to override auto-configuration (define your own bean; Spring backs off)
- [ ] profiles — how to configure per-environment

---

## Dependency Injection / IoC

**Key Terms:**
- `IoC (Inversion of Control)` — framework creates and wires objects instead of you
- `DI (Dependency Injection)` — IoC pattern; dependencies injected rather than created
- `bean` — object managed by Spring's IoC container
- `@Component` — generic bean; Spring detects and registers it
- `@Service` — @Component for service layer (semantic only; same as @Component)
- `@Repository` — @Component for data layer; adds exception translation
- `@Controller` — @Component for web layer
- `@Configuration` — class that defines beans with @Bean methods
- `@Bean` — method that returns an object to be registered as bean
- `@Autowired` — inject dependency; on constructor (preferred), field, or setter
- `constructor injection` — preferred; fields can be final; easier to test (no Spring needed in test)
- `field injection` — `@Autowired private Foo foo`; discouraged (harder to test, hides dependencies)
- `@Qualifier("name")` — when multiple beans of same type; specify which one to inject
- `@Primary` — marks bean as preferred when multiple exist
- `@Lazy` — initialize bean only when first requested
- `scope` — `singleton` (default; one per context), `prototype` (new per injection), `request`, `session`
- `bean lifecycle` — instantiate → inject dependencies → @PostConstruct → use → @PreDestroy → destroy

**Must understand:**
- [ ] Constructor injection vs field injection — why constructor is preferred
- [ ] Singleton scope — most beans are singletons; must be thread-safe
- [ ] @Bean vs @Component — explicit declaration vs classpath scanning
- [ ] How to test without Spring (inject through constructor; no @Autowired in test)

---

## Spring MVC / REST

**Key Terms:**
- `@RestController` — @Controller + @ResponseBody; all methods return response body directly
- `@RequestMapping("/path")` — map HTTP request to method; class-level prefix + method-level path
- `@GetMapping`, `@PostMapping`, `@PutMapping`, `@DeleteMapping`, `@PatchMapping` — shortcuts
- `@PathVariable` — extract from URL: `@GetMapping("/users/{id}") … @PathVariable Long id`
- `@RequestParam` — query string: `?page=1` → `@RequestParam(defaultValue="1") int page`
- `@RequestBody` — deserialize JSON request body to object (Jackson)
- `@ResponseBody` — serialize return value to JSON response
- `ResponseEntity<T>` — return type with control over status code and headers
- `@ResponseStatus` — set status code for method or exception
- `@ExceptionHandler` — handle exceptions in controller; return appropriate error response
- `@ControllerAdvice / @RestControllerAdvice` — global exception handler across all controllers
- `HttpStatus` — enum of HTTP status codes
- `Jackson` — JSON library; auto-configured by Spring Boot; @JsonProperty, @JsonIgnore
- `validation` — `@Valid` on @RequestBody; `@NotNull`, `@Min`, `@Max`, `@Email` on fields; returns 400 on failure
- `BindingResult` — capture validation errors without exception

**Error handling pattern:**
```java
@RestControllerAdvice
public class GlobalExceptionHandler {
    @ExceptionHandler(NotFoundException.class)
    @ResponseStatus(HttpStatus.NOT_FOUND)
    public ErrorResponse handleNotFound(NotFoundException ex) {
        return new ErrorResponse(ex.getMessage());
    }
}
```

**Must understand:**
- [ ] Difference between @Controller and @RestController
- [ ] How to return different HTTP status codes (ResponseEntity vs @ResponseStatus)
- [ ] Global exception handling with @RestControllerAdvice
- [ ] Request validation flow — @Valid + ConstraintViolationException → 400

---

## Spring Data JPA / Hibernate

**Key Terms:**
- `JPA (Java Persistence API)` — specification for ORM in Java
- `Hibernate` — most popular JPA implementation
- `@Entity` — marks class as persistent entity (mapped to DB table)
- `@Table(name="...")` — customize table name
- `@Id` — marks primary key field
- `@GeneratedValue` — auto-generate PK (IDENTITY, SEQUENCE, AUTO)
- `@Column` — customize column mapping (nullable, unique, name, length)
- `@OneToMany`, `@ManyToOne`, `@OneToOne`, `@ManyToMany` — relationship annotations
- `FetchType.LAZY` — load related entities only when accessed (default for collections)
- `FetchType.EAGER` — load related entities immediately with main entity
- `JPA Repository` — interface; Spring Data generates implementation; `CrudRepository`, `JpaRepository`
- `@Query` — custom JPQL or native SQL in repository
- `JPQL` — JPA Query Language; object-oriented SQL (refers to entity classes, not tables)
- `EntityManager` — JPA API for managing entity lifecycle (persist, merge, find, remove)
- `persistence context` — first-level cache; EntityManager's cache of managed entities
- `N+1 problem` — fetching list of entities, then Hibernate makes N queries for lazy relations; fix: JOIN FETCH or @BatchSize
- `JOIN FETCH` — JPQL: fetch association in one query
- `@Transactional` — method runs in transaction; commit on success, rollback on exception
- `dirty checking` — Hibernate automatically detects and persists changes to managed entities
- `detached entity` — entity not in persistence context; changes not automatically persisted

**Must understand:**
- [ ] N+1 problem — how to detect and fix (JOIN FETCH, @EntityGraph, @BatchSize)
- [ ] LAZY vs EAGER fetch — why LAZY is safer as default
- [ ] @Transactional — what it does and what exception triggers rollback (RuntimeException by default)
- [ ] Persistence context / first-level cache — why EntityManager is not thread-safe (request-scoped)

---

## Spring Security

**Key Terms:**
- `filter chain` — Spring Security is a chain of servlet filters; each request passes through
- `SecurityFilterChain` — configures which filters apply and in what order
- `authentication` — verifying who the user is (credentials check)
- `authorization` — verifying what the user can do (permission check)
- `SecurityContext` — holds authenticated user's `Authentication` object
- `SecurityContextHolder` — ThreadLocal holder of SecurityContext; access with `getContext()`
- `Authentication` — represents authenticated principal + credentials + authorities
- `UserDetails` — Spring's user representation; `loadUserByUsername()` in UserDetailsService
- `UserDetailsService` — interface; implement to load users from DB
- `PasswordEncoder` — `BCryptPasswordEncoder`; always hash passwords, never store plaintext
- `JWT filter` — custom filter that validates JWT and sets SecurityContext
- `OncePerRequestFilter` — base class for single-execution filters
- `@PreAuthorize("hasRole('ADMIN')")` — method-level security
- `@Secured("ROLE_ADMIN")` — older annotation for role check
- `CSRF protection` — Cross-Site Request Forgery token; usually disabled for stateless REST APIs
- `CORS` — Cross-Origin Resource Sharing; configure allowed origins for browser clients
- `OAuth2` — Spring Security supports OAuth2 client and resource server
- `@EnableWebSecurity` — enable Spring Security web security

**Must understand:**
- [ ] Filter chain — how authentication filters set SecurityContext
- [ ] Stateless JWT flow — filter validates token, sets Authentication, endpoint reads it
- [ ] When to disable CSRF (REST APIs with JWT; no session cookies)
- [ ] Role vs Authority — ROLE_ prefix convention

---

## Spring Transactions

**Key Terms:**
- `@Transactional` — Spring wraps method in a transaction (via AOP proxy)
- `propagation` — how transactions relate to existing ones:
  - `REQUIRED` (default) — join existing or create new
  - `REQUIRES_NEW` — always create new; suspend existing
  - `NESTED` — nested in existing transaction
  - `NOT_SUPPORTED` — suspend existing, run without transaction
  - `MANDATORY` — must have existing transaction; error if none
  - `NEVER` — must not have transaction; error if one exists
- `isolation levels` — READ_COMMITTED (default in Spring), READ_UNCOMMITTED, REPEATABLE_READ, SERIALIZABLE
- `rollback` — automatic on RuntimeException and Error; NOT on checked exceptions by default
- `rollbackFor = Exception.class` — rollback on checked exceptions too
- `noRollbackFor` — don't rollback on specific exception
- `@Transactional self-invocation` — calling @Transactional method from same class skips proxy! → transaction doesn't start; common bug
- `readOnly = true` — hint for optimization; prevents dirty checking; use on read-only methods
- `transaction timeout` — `@Transactional(timeout = 30)`

**Must understand:**
- [ ] Self-invocation problem — why `this.transactionalMethod()` skips the transaction
- [ ] REQUIRES vs REQUIRES_NEW — when each is needed
- [ ] Why checked exceptions don't trigger rollback by default
- [ ] readOnly transactions — what optimization it enables in Hibernate

---

## Kotlin Basics

**Key Terms:**
- `val` — immutable reference (like Java `final`)
- `var` — mutable reference
- `nullable types` — `String?` can be null; `String` cannot
- `safe call ?.` — `user?.address?.street` — returns null instead of NPE
- `elvis operator ?:` — `name ?: "default"` — use right side if left is null
- `non-null assertion !!` — throws NullPointerException if null (use rarely)
- `data class` — auto-generates equals, hashCode, toString, copy, componentN
- `sealed class` — restricted hierarchy; all subclasses in same file; great for when expressions
- `object` — singleton; Kotlin's replacement for static
- `companion object` — static members inside a class
- `extension function` — `fun String.trim(): String` — add methods to existing classes
- `lambda` — `{ x: Int -> x * 2 }`; last parameter convention: `list.map { it * 2 }`
- `it` — implicit parameter name in single-parameter lambda
- `scope functions` — `let`, `run`, `with`, `apply`, `also`
- `apply` — configure object: `Builder().apply { field = value }; return this`
- `let` — transform or null-check: `nullable?.let { process(it) }`
- `when` — powerful switch; exhaustive for sealed classes
- `data class copy` — `user.copy(name = "new name")` — immutable update

**Must understand:**
- [ ] Null safety — ? vs !! and why Kotlin eliminates most NPEs
- [ ] data class vs regular class — what auto-generated methods give you
- [ ] Extension functions — how they work (compiled as static functions with receiver)
- [ ] sealed class + when — exhaustive pattern matching without else

---

## Kotlin Coroutines

**Key Terms:**
- `coroutine` — suspendable computation; like goroutine but cooperative
- `suspend` — function that can be paused without blocking a thread
- `CoroutineScope` — defines lifecycle of coroutines; cancel scope = cancel all children
- `launch` — fire-and-forget coroutine; returns `Job`
- `async` — coroutine that returns `Deferred<T>`; `await()` gets result
- `runBlocking` — bridge blocking code to coroutines; use in tests and main
- `Dispatchers.IO` — for blocking I/O (file, network); uses thread pool
- `Dispatchers.Default` — for CPU-intensive work; uses CPU core count threads
- `Dispatchers.Main` — for Android UI thread (N/A for backend)
- `withContext(Dispatchers.IO)` — switch dispatcher for a block; suspend
- `Job` — handle for coroutine; `cancel()`, `join()`
- `SupervisorJob` — child failures don't propagate to parent or siblings
- `Flow` — cold asynchronous stream; like RxJava Observable but simpler
- `StateFlow` — hot flow; always has value; like LiveData
- `SharedFlow` — hot flow for events; multiple collectors
- `collect` — terminal operator; starts collecting a Flow
- `structured concurrency` — coroutines have parents; cancellation propagates down; errors propagate up

**Must understand:**
- [ ] suspend function — what it means (can be paused; not blocking a thread)
- [ ] launch vs async — fire-and-forget vs need result
- [ ] Dispatchers — which to use for I/O vs CPU work
- [ ] Structured concurrency — why all coroutines should be in a scope (leak prevention)

---

## Maven / Gradle

**Key Terms:**
- `POM (Project Object Model)` — Maven's `pom.xml`; declares dependencies, plugins, build lifecycle
- `groupId:artifactId:version` — Maven coordinate uniquely identifies a dependency
- `local repository` — `~/.m2/repository`; caches downloaded dependencies
- `central repository` — Maven Central; default source for dependencies
- `dependency scope` — `compile` (default), `test`, `runtime`, `provided`
- `build lifecycle` — phases: validate → compile → test → package → install → deploy
- `plugin` — extends Maven; `maven-compiler-plugin`, `maven-surefire-plugin` (tests)
- `multi-module` — parent POM with `<modules>`; child modules inherit parent config
- `BOM (Bill of Materials)` — dependency management POM; standardizes versions (Spring Boot BOM)
- `Gradle` — Groovy/Kotlin DSL build tool; more flexible than Maven; faster incremental builds
- `build.gradle.kts` — Kotlin DSL build script (preferred over Groovy)
- `Gradle tasks` — `./gradlew build`, `./gradlew test`, `./gradlew bootRun`
- `Gradle dependency configuration` — `implementation` (compile), `testImplementation`, `runtimeOnly`
- `Gradle wrapper` — `gradlew`; pins Gradle version for reproducible builds
- `dependency resolution` — transitive dependencies; version conflicts resolved by nearest definition (Maven) or latest (Gradle)

**Must understand:**
- [ ] Maven lifecycle phases — what each does
- [ ] `implementation` vs `api` in Gradle — api exposes dependency to consumers; implementation doesn't
- [ ] BOM — why it's useful (don't manage individual Spring dependency versions)
- [ ] How to exclude a transitive dependency (when version conflict causes issues)
