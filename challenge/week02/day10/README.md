# Day 10 of #66DaysOfGo

_Last update:  Jul 22, 2023_.

---

Today, I've continued with the WebAssembly series, this time reading an article written in 2018 about "WebAssembly Future Skills".

---

## WebAssembly Future Skills

> _Based on [https://hacks.mozilla.org/2018/10/webassemblys-post-mvp-future/](https://hacks.mozilla.org/2018/10/webassemblys-post-mvp-future/)_ by [Lin Clark](https://twitter.com/linclark), [Till Schneidereit](https://hacks.mozilla.org/author/tschneidereitmozilla-com/), [Luke Wagner](https://hacks.mozilla.org/author/lwagnermozilla-com/).

<img src="https://hacks.mozilla.org/files/2018/10/01-07-runtime-09-final-e1539904436477.png" alt="WebAssembly’s skill tree" width="350"/>

Summary:

This article discusses the misconception that the WebAssembly introduced in browsers in 2017 is its final version, due to its backward compatibility. However, WebAssembly is not feature-complete, with many features still to be introduced, which will significantly enhance its capabilities.

WebAssembly began with [Emscripten](https://emscripten.org/), allowing C++ code to run on the web by _transpiling_ (converting source code from a programming language into an equivalent source code of the same or a different programming language) it to JavaScript. This made it possible to bring large C++ codebases to the web. Later, a faster JavaScript subset named `asm.js` was developed.

WebAssembly was created for even faster performance. The first version of WebAssembly required a language-agnostic compile target, fast execution, compactness, and linear memory.

WebAssembly's first achievements were to enable desktop applications and games to run efficiently in browsers. The next goal is to allow heavy-weight desktop applications to run in browsers. Examples of these already include Autodesk’s AutoCAD and Adobe's Lightroom. However, additional features are needed to ensure optimal performance of all applications in the browser.

It also discusses several future enhancements for WebAssembly to improve its performance and capabilities:

- **Threading**: Necessary for utilizing multiple cores in modern computers. As of 2018, a proposal was nearly complete ([https://web.dev/webassembly-threads/](https://web.dev/webassembly-threads/)).
- **SIMD** (Single Instruction Multiple Data): Allows for parallel processing, currently under active development.
- **64-bit Addressing**: WebAssembly needs this to remove the artificial limit on address space, plans for its integration are underway.
- **Streaming Compilation**: Allows compiling a WebAssembly file while it's being downloaded, already added in Firefox.
- **Tiered Compilation**: Involves using two compilers for quicker start-up times and faster execution times. A new optimizing compiler, [Cranelift](https://cranelift.dev/), is being developed to improve performance further.
- **Implicit HTTP Caching**: Enables storing of compiled code in the HTTP cache, thereby skipping the need for re-compiling for repeat page visits.
- **Fast calls between JS and WebAssembly**: Initially, these calls were slow but improvements have been made in Firefox and other browsers.
- **Easy and fast data exchange**: Currently, data passing is slow and difficult, requiring conversion of complex objects into numbers. Proposals exist to facilitate this process.
- **ES module integration**: Currently, WebAssembly modules are not part of the JS module graph. This needs to change to allow for importing and exporting like JS modules.
- **Toolchain integration**: WebAssembly needs a place for module distribution and bundling tools like npm or webpack.
- **Backwards compatibility**: Support for older browsers, including those that don't recognize WebAssembly, is essential.
- **GC integration with the browser**: It helps in optimizing JavaScript frameworks by rewriting parts using other languages for parallel processing. It also aids in efficient memory allocation, reducing the need for creating objects for garbage collection. This integration supports cross-language data dependencies and is beneficial for statically-typed languages that compile to JS.
- **Exception handling**: Necessary since some languages extensively use it, and the current _polyfill_ (A piece of code, usually JavaScript on the Web, used to provide modern functionality on older browsers that do not natively support it) method slows down code execution. Improved support will help WebAssembly correctly handle exceptions thrown by JavaScript functions.
- **Improved debugging support for WebAssembly in browsers**.
- **Support for tail calls**: Necessary for many functional languages, allowing the calling of a new function without adding a new stack frame to the stack.

NodeJS:

- WebAssembly could significantly increase Node.js's portability by serving as an alternative to Node's native modules, which need to be compiled for each specific machine. WebAssembly modules would not need such specific compilation and could run at nearly native performance. However, WebAssembly does not have direct system resource access as native modules do. Thus, to enable WebAssembly to work effectively with Node.js, functions allowing WebAssembly modules to interact with the operating system must be included.
- There's a possibility of creating a standardized set of functions, similar to a POSIX for WebAssembly, usable across different runtime environments. This could lead to a unified API for calling functions, regardless of the platform, facilitating universal modules' creation.
- The "package name maps" proposal is a step towards this direction, aiming to map module names to paths, thus allowing browsers and Node to load different modules with the same API. Despite no active work on the rest of the ideas, there is strong discussion suggesting potential implementation.
- WebAssembly's use could extend beyond browsers to other applications, including CDNs, serverless, and edge computing. These use cases could benefit from WebAssembly's safety, speed, and scale, as demonstrated by Fastly, a CDN and edge computing provider. Fastly needed to create its own runtime to use WebAssembly, involving compiling WebAssembly down to machine code and creating custom functions to interact with the system. Using a common runtime across different companies could streamline development significantly.

Even though no standard runtime exists yet, several projects are underway, like WAVM, wasmjit, and wasmtime, which is based on the Cranelift compiler used by Fastly. A common runtime can speed up the development for various use cases.
WebAssembly can also be used in more traditional operating systems, enabling portable command-line tools across different systems. Moreover, it could be advantageous for Internet of Things (IoT) devices, which are typically resource-constrained. WebAssembly's portability and low-memory footprint would be beneficial here.

## References

- [WebAssembly’s post-MVP future: A cartoon skill tree](https://hacks.mozilla.org/2018/10/webassemblys-post-mvp-future/)
- [https://hacks.mozilla.org/category/code-cartoons/](https://hacks.mozilla.org/category/code-cartoons/)
- [https://web.dev/webassembly-threads/](https://web.dev/webassembly-threads/)
- [(Github) WebAssembly proposals](https://github.com/WebAssembly/proposals)
