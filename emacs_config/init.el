; On a Mac, how to get the meta key to work:
; http://stackoverflow.com/questions/162896/emacs-on-mac-os-x-leopard-key-bindin
gs

(add-to-list 'load-path "~/workspace/git/go/emacs_config/")
(require 'go-mode-autoloads)

(add-hook 'before-save-hook 'gofmt-before-save)
