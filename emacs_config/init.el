(add-to-list 'load-path "~/workspace/git/go/emacs_config/")
(require 'go-mode-autoloads)

(add-hook 'before-save-hook 'gofmt-before-save)
