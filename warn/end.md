# Заключение

Недавно я услышал о том, что Go – это скучный язык. Скучный потому, что его легко изучить, легко на нём писать и, что самое главное, легко читать. Возможно, я оказал вам медвежью услугу. Мы потратили три главы на разговоры о типах и о том, как объявить переменную.

Если у вас уже был опыт работы со статически типизированным языком, многое из того, что вы видели, в лучшем случае освежило ваши знания. То, что Go делает видимыми указатели и то, что срезы являются просто тонкими обёртками вокруг массивов, скорее всего, не удивительно для опытных Java или C# разработчиков.

Если вы в основном использовали динамические языки, вы могли почувствовать небольшую разницу. Её нужно понять. Не на последнем месте стоит также синтаксис объявления и инициализации. Несмотря на то, что я фанат Go, я считаю, что весь успех в достижении простоты. И все сводится к простым правилам (например то, что вы можете объявить переменную только один раз и := делает это) и фундаментальным понятиям (то, что new(X) или &X{} только выделяет память, но срезы, карты и каналы требуют дополнительных действий при инициализации и поэтому нужен make).

Помимо этого, Go предоставляет простой, но эффективный способ организации кода. Интерфейсы, основанная на возвращении обработка ошибок, defer для управления ресурсами и простой способ достижения композиции.

Последним, но важным, является встроенная поддержка конкурентности. Горутины довольно эффективны и просты (просты в использовании по крайней мере). Это хорошая абстракция. Каналы немного сложнее. Я всегда считал, что понимание базовых вещей важно перед использованием высокоуровневых обёрток. Я думаю полезно узнать о конкурентном программировании без каналов. Всё же каналы реализованы таким образом, что они совсем не похожи на простую абстракцию. Они почти как самостоятельные фундаментальные блоки. Я говорю это потому, что они изменяют стиль написания и понимания конкурентного программирования. Учитывая то, каким сложным конкурентное программирование может быть, реализация определённо хороша.
