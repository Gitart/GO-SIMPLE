### Способы визуализации данных профилирования в Golang

Инструменты Go обеспечивают визуализацию данных профиля с помощью text, graph и callgrind визуализации с помощью go tool pprof.

[![](https://2.bp.blogspot.com/-NpldfrCCxsk/Xo3_fqJU-KI/AAAAAAAABVE/Usp9b380RxYNggeXFWuqr7KXw_QCHHVaACPcBGAYYCw/s400/viz.png)](https://2.bp.blogspot.com/-NpldfrCCxsk/Xo3_fqJU-KI/AAAAAAAABVE/Usp9b380RxYNggeXFWuqr7KXw_QCHHVaACPcBGAYYCw/s1600/viz.png)

Распечатка самых дорогих вызовов в виде text.

[![](https://4.bp.blogspot.com/-8oMMwi91-BE/Xo3_2Ouoy_I/AAAAAAAABVM/kAKmhVfEX7QU9MH0d8hYU2CksY3rNgIVQCLcBGAsYHQ/s400/viz2.png)](https://4.bp.blogspot.com/-8oMMwi91-BE/Xo3_2Ouoy_I/AAAAAAAABVM/kAKmhVfEX7QU9MH0d8hYU2CksY3rNgIVQCLcBGAsYHQ/s1600/viz2.png)

Визуализация самых дорогих вызовов в виде graph.

Представление веб\-списка (weblist) отображает дорогие части исходного кода построчно на странице HTML. В следующем примере 530мс тратится в runtime.concatstrings, а стоимость каждой строки представлена в листинге.

[![](https://4.bp.blogspot.com/-JX6QLPqB81g/Xo3_77rLyZI/AAAAAAAABVQ/Cp8n_Ybd7R4oLd2rv84cTaRSZr8h6I57wCLcBGAsYHQ/s400/viz3.png)](https://4.bp.blogspot.com/-JX6QLPqB81g/Xo3_77rLyZI/AAAAAAAABVQ/Cp8n_Ybd7R4oLd2rv84cTaRSZr8h6I57wCLcBGAsYHQ/s1600/viz3.png)

Визуализация самых дорогих звонков в виде веб\-списка (weblist).

Другим способом визуализации данных профиля является flame graph (график пламени). flame graph позволяют вам перемещаться по определенному пути предков, так что вы можете увеличивать/уменьшать отдельные участки кода. В upstream pprof есть поддержка flame graph.

[![](https://2.bp.blogspot.com/-lunJfdD12p8/Xo4ABZ2NI1I/AAAAAAAABVU/x-wH4aIXjOgmrQgGp2KXgq52DYXiMC9sQCLcBGAsYHQ/s400/viz4.png)](https://2.bp.blogspot.com/-lunJfdD12p8/Xo4ABZ2NI1I/AAAAAAAABVU/x-wH4aIXjOgmrQgGp2KXgq52DYXiMC9sQCLcBGAsYHQ/s1600/viz4.png)

flame graph предлагают визуализацию для определения самых дорогих путей кода.
