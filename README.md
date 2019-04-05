# zxp golang
# windows serial debug tools
# depend module
# github.com/lxn/walk
# please move serialcontrol files in folder of serialcontrol to walk folder
usage : go get -u -v github.com/pzhu1975/zxpgolang

Go语言有一个不错的设计，就是build constraints(构建约束)。可以在源码中通过注释的方式指定编译选项，比如只允许在linux下，或者在386的平台上编译啊之类的；还可以通过文件名来约束构造，比如xxxx_linux.go，就是只允许在linux下编译，xxx_windows_amd64.go就是只允许在windows x64下编译。
构建约束可以在很多文件中使用，不单单是GO文件。但是必须要注意的是，通过注释实施构建约束的话，比如要放在文件的开头，要优先于空行或和其他注释之前。也就是说必须在package语句的前面写。这就有个很蛋疼的文件，因为GO的godoc是可以提取代码中的注释然后转换为文档的。在package语句之前写的注释会被认为是包级别的注释。而构建约束又在所有注释之前，那么为了区分包级别的注释，就要在构建约束与包级别的注释之间添加空行进行区分。（这个设计看上去不得不承认很囧）。

通过注释实施的构建约束还可以进行逻辑表达。就是and, or之类的语义。GO的官方是这么定义的：如果构建约束中有空格，那么就是OR关系，如果是逗号分隔，那么就是AND关系。！表示not。比如
//+build linux,386 darwin,!cgo

就是表示(linux AND 386) OR (darwin AND (NOT cgo))

而且GO还支持多行的构建约束，多行之间是AND关系，比如

// +build linux darwin

// +build 386

就是表示(linux OR darwin) AND 386

GO官方还定义了常用的一些约束

限制目标操作系统，也就是要和runtime.GOOS一致
限制目标架构平台，也就是要和runtime.GOARCH一致
GC或者GCCGO等编译器支持的约束
cgo约束，也就是说如果支持cgo的话，就可以参与编译
go1.1，表示从go1.1开始向前兼容

如果你想临时让某个文件不参与编译，可以添加注释约束下： // +build ignore
通过注释来实现构建约束有点蛋疼，而且GO官方定义里还表示可以自定义约束，那么可以用来干嘛？学GO的人都知道GO内建了单元测试的框架，跑跑一般的单元测试还是非常嗨皮的，但是如果要做一些简单的集成测试就令人拙计了，因为go test默认就是跑最基本的单元测试。那么怎么只执行集成测试的代码呢？其实就可以通过构建约束来实施。比如我们在集成测试的GO文件中加上 // +build integration 然后运行命令 go test –tags=”integration”就可以只运行我们的集成测试代码了。

虽然通过注释的方式对构建进行了约束，但是文件名的构建约束反而让人觉着很不错呢，至少在工程里看着一目了然。文件名前缀只要含有_GOOS, _GOARCH, _GOOS_GOARCH的就可以了。比如xxx_linux.go yyy_windows_amd64.go ， zzzz_386.s等等。
