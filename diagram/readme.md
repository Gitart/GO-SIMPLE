## Plan and Craph

https://mermaid-js.github.io/mermaid/#/gantt
https://mermaid-js.github.io/mermaid/#/Tutorials
https://docs.github.com/en/get-started/writing-on-github/working-with-advanced-formatting/creating-diagrams



```mermaid
gantt
dateFormat  YYYY-MM-DD
title Adding GANTT diagram to mermaid
excludes weekdays 2014-01-10

section A section
Completed task            :done,    des1, 2014-01-06,2014-01-08
Active task               :active,  des2, 2014-01-09, 3d
Future task               :         des3, after des2, 5d
Future task2               :         des4, after des3, 5d
```

## Mer
```mermaid
sequenceDiagram
    participant Alice
    participant Bob
    Alice->>John: Hello John, how are you?
    loop Healthcheck
        John->>John: Fight against hypochondria
    end
    Note right of John: Rational thoughts <br/>prevail!
    John-->>Alice: Great!
    John->>Bob: How about you?
    Bob-->>John: Jolly good!
```


## Diagram
```mermaid
sequenceDiagram
    par Alice to Bob
        Alice->>Bob: Hello guys!
    and Alice to John
        Alice->>John: Hello guys!
    end
    Bob-->>Alice: Hi Alice!
    John-->>Alice: Hi Alice!
```

## Satate diagram

```mermaid
stateDiagram
    direction LR
    [*] --> A
    A --> B
    B --> C
    state B {
      direction LR
      a --> b
    }
    B --> D
```












https://mermaid.live/edit

[![](https://mermaid.ink/img/pako:eNpt0M2qAjEMBeBXidk68wJdKILC1a3bbkJ7xin051pbRMR3tzOOO7MK5DuB5MkmWbDiG64V0WDv5JIl6Eitdt4Z9JvN-pTGqOgP3iea-o7GdCfJoEeq25_4w4zEidAIyRSw0GnWN9rPmbbYfdIdHefErFts9ZsfaQA8XTKkLORb3HFADuJsu-k5zTSXEQGaVWstBqm-aNbx1Wj9t1JwsK6kzGoQf0PHUks6P6JhVXLFFy1_WdTrDbKeYug)](https://mermaid.live/edit#pako:eNpt0M2qAjEMBeBXidk68wJdKILC1a3bbkJ7xin051pbRMR3tzOOO7MK5DuB5MkmWbDiG64V0WDv5JIl6Eitdt4Z9JvN-pTGqOgP3iea-o7GdCfJoEeq25_4w4zEidAIyRSw0GnWN9rPmbbYfdIdHefErFts9ZsfaQA8XTKkLORb3HFADuJsu-k5zTSXEQGaVWstBqm-aNbx1Wj9t1JwsK6kzGoQf0PHUks6P6JhVXLFFy1_WdTrDbKeYug)



[![](https://mermaid.ink/img/pako:eNqNjzEOwjAMRa9See4JMoM4AGsWk7hJRBJHwUFCVe_egGgXEOJP9vf7X_IMhi2BgqHLBTlVLF7n4SXDKQX5vl0qZuMHS3eKXHbGk7lykw__R9OeSRjyZiaqjv5rgRE63cO2vzE_rxrEUyINqo-WJmxRNOi8dLQVi0JHG4QrqAnjjUbAJnx-ZANKaqMNOgR0FdObWlYJ4l-7)](https://mermaid.live/edit#pako:eNqNjzEOwjAMRa9See4JMoM4AGsWk7hJRBJHwUFCVe_egGgXEOJP9vf7X_IMhi2BgqHLBTlVLF7n4SXDKQX5vl0qZuMHS3eKXHbGk7lykw__R9OeSRjyZiaqjv5rgRE63cO2vzE_rxrEUyINqo-WJmxRNOi8dLQVi0JHG4QrqAnjjUbAJnx-ZANKaqMNOgR0FdObWlYJ4l-7)
