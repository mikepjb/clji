# clji - Clojure (Command Line) Interface

A tool used to send code between vim and a running nREPL.

Test with `lein repl :start :port 9999` (I have no idea where `.nrepl-port` went..)

## Use in Vim

```
func! Require()
  let call = system('clji "(load-file \"' . expand('%') . '\")"')
  echo call
endfunc

command! Require :call Require()
```

```
command! -nargs=? -range Eval <line1>,<line2>call Eval('<args>')
vnoremap <CR> :Eval<CR>
function! Eval(regex) range
  let section = getline(a:firstline, a:lastline)
  let call = system("clji \"" . join(section) . "\"")
  echo call
endfunction
```

## Notes on nREPL

if no project.clj is found, it creates a ~/.lein/nrepl-port file
otherwise .nrepl-port is used at the root of a project.

returns out key and value key
  - out: println
  - value: return value
  e.g (println 42), out: "42\n", value: nil

## Reference

https://andrewvos.com/writing-async-jobs-in-vim-8/
