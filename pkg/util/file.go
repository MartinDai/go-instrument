package util

import (
	"go-instrument/pkg/instrument"
	"go-instrument/pkg/logger"
	"go/ast"
	"go/format"
	"go/parser"
	"go/token"
	"os"
	"path/filepath"
)

// InstrumentFunc 根据函数名增强函数
func InstrumentFunc(name string) error {
	err := filepath.Walk(".", func(path string, info os.FileInfo, err error) error {
		if err != nil || info.IsDir() || filepath.Ext(path) != ".go" {
			return nil
		}

		//解析文件
		fileSet := token.NewFileSet()
		logger.Debug("parsing file(%v)", path)
		file, err := parser.ParseFile(fileSet, path, nil, parser.ParseComments)
		if err != nil {
			logger.Debug("Error parsing file(%v), cause:%v", path, err)
			return nil
		}

		// 匹配函数
		var instrumented = false
		ast.Inspect(file, func(n ast.Node) bool {
			fn, ok := n.(*ast.FuncDecl)
			if !ok {
				return true
			}
			if fn.Name.Name == name {
				logger.Info("Found need instrument func %v in(%v)", name, path)
				instrument.AddPrintln(fn)
				instrumented = true
				return false
			}
			return true
		})

		if instrumented {
			saveInstrumentedFile(path, fileSet, file)
		}

		return nil
	})

	return err
}

// 保存增强后的文件
func saveInstrumentedFile(path string, fileSet *token.FileSet, file *ast.File) {
	out, err := os.Create(path)
	if err != nil {
		logger.Error(err, "Error creating file:%v", path)
		return
	}

	if err := format.Node(out, fileSet, file); err != nil {
		logger.Error(err, "Save instrument file error, file path:%v", path)
	}
}
