[TOC]

# Add Font
https://github.com/lusingander/fyne-font-example
```
$ fyne bundle mplus-1c-regular.ttf > bundle.go
$ fyne bundle -append mplus-1c-bold.ttf >> bundle.go
```

# Package with icon
https://developer.fyne.io/started/packaging
https://developer.fyne.io/started/packaging-mobile
```
fyne package -os linux -icon myapp.png
fyne package -os windows -icon myapp.png
```

# Distribution
https://developer.fyne.io/started/distribution
```
macOS App Store since 1.4.2
Android         since 1.x.x
iOS App Store   since 1.4.1
```

# Compile
## Option
https://developer.fyne.io/started/compiling
```
構建標籤
Fyne 通常會通過選擇驅動程序和配置來為目標平台適當地配置您的應用程序。支持以下構建標記，可以幫助您進行開發。例如，如果您希望在台式計算機上運行時模擬移動應用程序，則可以使用以下命令：

go run -tags mobile main.go
標籤	    描述
gles	強制使用嵌入式 OpenGL (GLES) 而不是完整的 OpenGL。這通常由目標設備控制，通常不需要。
hints	顯示改進或優化的開發人員提示。hints當您的應用程序不遵循 Material Design 或其他建議時，Running with將記錄日誌。
mobile	此標記在模擬的移動窗口中運行應用程序。當您想在移動平台上預覽您的應用程序而無需編譯和安裝到設備上時很有用。
no_native_menus	此標誌專用於 macOS，指示應用程序不應使用 macOS 本機菜單。相反，菜單將顯示在應用程序窗口內。對於在 macOS 上測試應用程序以模擬 Windows 或 Linux 上的行為最有用。
```
## Cross Compile
https://developer.fyne.io/started/cross-compiling
```

```

---

# Others

https://gist.github.com/sighingnow/deee806603ec9274fd47
https://github.com/mouuff/go-rocket-update-example

---
# RW
```
https://juejin.cn/post/6864886461746855949
src:
  - bufio
    - bufio.go
  - bytes
    - buffer.go
    - reader.go
  - io
    - ioutil
      - ioutil.go
    - io.go
  - os
    - file.go
  - strings  
```