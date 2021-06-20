Товары и услуги

Раздел "Товары и услуги" содержит модели данных, обеспечивающие внешним приложениям доступ к каталогу товаров и услуг, который ведется в системе Бизнес.Ру.

Справочник товаров и услуг (goods)

**Описание модели:** Cправочник товаров и услуг

**URL ресурса:** https://myaccount.business.ru/api/rest/goods.json

**Разрешенные запросы:** get(чтение), post(создание), put(изменение), delete(удаление)

**Специальные параметры модели:**

**group\_ids** - массив id групп товаров (модель groupsofgoods). При передаче в GET запросе group\_ids в ответе будут возвращены товары в этих группах и во всех их дочерних группах.

**with\_attributes** - при передаче в GET запросе with\_attributes=1 в ответ будет добавлен массив атрибутов товара.

**with\_barcodes** - при передаче в GET запросе with\_barcodes =1 в ответ добавляется массив штрихкодов товара.

**with\_remains** - при передаче в GET запросе with\_remains=1 в ответ добавляется массив остатков товаров по складам, указанным в store\_ids. Если store\_ids не передается, то возвращается массив остатков товаров по всем складам. По каждому складу выводится общий остаток товара (total), а также зарезервированное на данном складе количество товара (reserved).

**store\_ids** - массив id складов (модель stores), по которым следует возвратить остатки. Используется совместно с with\_remains.

**filter\_positive\_remains** - при передаче в GET запросе filter\_positive\_remains =1 в ответе будут возвращены товары с положительным суммарным остатком по складам, указанным в store\_ids. Если store\_ids не передается, то будут возвращены товары с положительным суммарным остатком по всем складам. Используется совместно с with\_remains.

**with\_prices** - при передаче в GET запросе with\_prices =1 в ответ добавляется массив отпускных цен по типам, указанным в type\_price\_ids. Если type\_price\_ids не передается, то будут возвращены все отпускные цены.

**type\_price\_ids** - массив id типов цен (модель salepricetypes), по которым следует возвратить цены. Используется совместно с with\_prices.

**with\_modifications** - при передаче в GET запросе with\_modifications =1 в ответ добавляется массив модификаций товара. При одновременном указании параметров with\_attributes, with\_barcodes, with\_remains, with\_prices, соответствующая информация будет возвращена для каждой модификации товара.

| Параметр | Тип | Обязателен в запросах | Внешний ключ к модели | Значение по умолчанию | Описание |
| --- | --- | --- | --- | --- | --- |
| id | int | put, delete |  |  | Идентификатор |
| name | string | post |  |  | Наименование товара |
| full\_name | string |  |  |  | Полное наименование товара, используемое в печатных документах |
| nds\_id | int |  | [nds](https://developers.business.ru/api-polnoe/stavki_nds_nds/361) |  | Ссылка на НДС |
| group\_id | int |  | [groupsofgoods](https://developers.business.ru/api-polnoe/derevo_grupp_groupsofgoods/290) |  | Ссылка на группу товаров |
| part | string |  |  |  | Артикул |
| store\_code | string |  |  |  | Код товара на складе |
| type | int |  |  |  | Тип позиции (1-товар, 2-услуга, 3-комплект) |
| archive | bool |  |  |  | Флаг (0-товар не перемещен в архив, 1-товар перемещен в архив) |
| description | string |  |  |  | Описание товара |
| country\_id | int |  | [countries](https://developers.business.ru/api-polnoe/strany_countries/355) |  | Ссылка на страну производителя |
| allow\_serialnumbers | bool |  |  |  | Флаг (0-учет по серийным номерам запрещен, 1-учет по серийным номерам разрешен) |
| weight | float |  |  |  | Вес товара |
| volume | float |  |  |  | Объем товара |
| code | string |  |  |  | Идентификатор товара |
| store\_box | string |  |  |  | Код ячейки на складе |
| remains\_min | float |  |  |  | Минимально допустимый остаток товара на складе |
| partner\_id | int |  | [partners](https://developers.business.ru/api-polnoe/spravochnik_kontragentov_partners/303) |  | Ссылка на поставщика |
| responsible\_employee\_id | int |  | [employees](https://developers.business.ru/api-polnoe/spisok_sotrudnikov_employees/343) |  | Ссылка на сотрудника, ответственного за товар |
| images | images |  |  |  | Изображения товара |
| feature | int |  |  |  | Особенность учёта: 1-весовой, 2-Алкоголь, 3-Пиво, 4-Маркируемый |
| departments\_ids | int\[\] |  |  |  | Массив ссылок на отделы |
| cost | float |  |  |  | Себестоимость |
| measure\_id | int |  | [measures](https://developers.business.ru/api-polnoe/edinitsy_izmereniya_measures/349) |  | Ссылка на основную еденицу измерения товара |
| good\_type\_code | string |  |  |  | Код вида товара |
| payment\_subject\_sign | int |  |  |  | Признак предмета расчета (1-Товар, 2-Подакцизный товар, 3-Работа, 4-Услуга, 7-Лотерейный билет, 12-Составной предмет расчета, 13-Иной предмет расчета) |
| marking\_type | int |  |  |  | Вид маркир. товара: 1 - Табачная продукция, 2 - Шубы, 3 - Обувь, 4 - Лекарства, 5 - Молочная продукция, 6 - Кресла-коляски, 7 - Велосипеды, 8 - Фотоаппараты, 9 - Шины и покрышки, 10 - Легкая промышленность, 11 - Духи и туалетная вода, 12 - Средства индивидуальной защиты |
| allow\_marking | boolean |  |  |  | Признак маркируемого товара. 1-Маркируемый товар, 0-Немаркируемый товар |
| taxation | int |  |  |  | Система налогообложения (для розницы). 0-ОСН, 1-УСН доход, 2-УСН доход - расход, 3-ЕНВД, 4-ЕСН, 5-Патент |
| updated | datetime |  |  |  | Время последнего обновления |
| deleted | bool |  |  |  | Строка удалена (перемещена в корзину) |

[Подробнее о запросах](https://developers.business.ru/api-polnoe/zaprosy_k_api/369)[Примеры запросов](https://developers.business.ru/api-polnoe/primery_zaprosov_k_api/377)[О дополнительных параметрах](https://developers.business.ru/api-polnoe/poluchenie_dannyh_iz_svyazannyh_modelej/454)Сформировать запрос к модели

Дерево групп (groupsofgoods)

**Описание модели:** Дерево групп товаров и услуг (рубрикатор)

**URL ресурса:** https://myaccount.business.ru/api/rest/groupsofgoods.json

**Разрешенные запросы:** get(чтение), post(создание), put(изменение), delete(удаление)

**Специальные параметры модели:**

**group\_ids** - массив id групп товаров. При передаче в GET запросе group\_ids в ответе будут возвращены ветви дерева групп, в пределах которых расположены заданные группы.

| Параметр | Тип | Обязателен в запросах | Внешний ключ к модели | Значение по умолчанию | Описание |
| --- | --- | --- | --- | --- | --- |
| id | int | put, delete |  |  | Идентификатор |
| name | string | post |  |  | Наименование группы |
| parent\_id | int |  | [groupsofgoods](https://developers.business.ru/api-polnoe/derevo_grupp_groupsofgoods/290) |  | Ссылка на родительскую группу |
| images | images |  |  |  | Изображения группы товаров |
| updated | datetime |  |  |  | Время последнего обновления |
| deleted | bool |  |  |  | Строка удалена (перемещена в корзину) |

[Подробнее о запросах](https://developers.business.ru/api-polnoe/zaprosy_k_api/369)[Примеры запросов](https://developers.business.ru/api-polnoe/primery_zaprosov_k_api/377)[О дополнительных параметрах](https://developers.business.ru/api-polnoe/poluchenie_dannyh_iz_svyazannyh_modelej/454)Сформировать запрос к модели

Привязка единиц измерения (goodsmeasures)

**Описание модели:** Привязка единиц измерения к товарам

**URL ресурса:** https://myaccount.business.ru/api/rest/goodsmeasures.json

**Разрешенные запросы:** get(чтение), post(создание), put(изменение), delete(удаление)

| Параметр | Тип | Обязателен в запросах | Внешний ключ к модели | Значение по умолчанию | Описание |
| --- | --- | --- | --- | --- | --- |
| id | int | put, delete |  |  | Идентификатор |
| good\_id | int | post |  |  | Ссылка на товар |
| measure\_id | int | post |  |  | Ссылка на единицу измерения |
| base | bool |  |  |  | Флаг( 0-небазовая единица измерения, 1-базовая единица измерения ) |
| coefficient | float |  |  |  | Коеффициент единицы измерения относительно базовой |

[Подробнее о запросах](https://developers.business.ru/api-polnoe/zaprosy_k_api/369)[Примеры запросов](https://developers.business.ru/api-polnoe/primery_zaprosov_k_api/377)[О дополнительных параметрах](https://developers.business.ru/api-polnoe/poluchenie_dannyh_iz_svyazannyh_modelej/454)Сформировать запрос к модели

Штрихкоды (barcodes)

**Описание модели:** Штрихкоды товаров и модификаций товаров

**URL ресурса:** https://myaccount.business.ru/api/rest/barcodes.json

**Разрешенные запросы:** get(чтение), post(создание), put(изменение), delete(удаление)

| Параметр | Тип | Обязателен в запросах | Внешний ключ к модели | Значение по умолчанию | Описание |
| --- | --- | --- | --- | --- | --- |
| id | int | put, delete |  |  | Идентификатор |
| good\_id | int |  | [goods](https://developers.business.ru/api-polnoe/spravochnik_tovarov_i_uslug_goods/289) |  | Ссылка на товар |
| modification\_id | int |  | [goodsmodifications](https://developers.business.ru/api-polnoe/privyazka_modifikatsij_goodsmodifications/301) |  | Ссылка на модификацию товара |
| type | int |  |  |  | Тип штрихкода |
| value | string | post |  |  | Значение штрихкода |
| updated | datetime |  |  |  | Время последнего обновления |
| deleted | bool |  |  |  | Строка удалена (перемещена в корзину) |

[Подробнее о запросах](https://developers.business.ru/api-polnoe/zaprosy_k_api/369)[Примеры запросов](https://developers.business.ru/api-polnoe/primery_zaprosov_k_api/377)[О дополнительных параметрах](https://developers.business.ru/api-polnoe/poluchenie_dannyh_iz_svyazannyh_modelej/454)Сформировать запрос к модели

Атрибуты (характеристики) товаров

Данный подраздел содержит модели данных, обеспечивающие внешним приложениям доступ к информации об атрибутах товаров и услуг.

Атрибуты (attributesforgoods)

**Описание модели:** Список атрибутов товаров и услуг

**URL ресурса:** https://myaccount.business.ru/api/rest/attributesforgoods.json

**Разрешенные запросы:** get(чтение), post(создание), put(изменение), delete(удаление)

| Параметр | Тип | Обязателен в запросах | Внешний ключ к модели | Значение по умолчанию | Описание |
| --- | --- | --- | --- | --- | --- |
| id | int | put, delete |  |  | Идентификатор |
| name | string | post |  |  | Наименование атрибута |
| selectable | bool | post |  |  | Флаг, определяющий имеет ли атрибут предопределенный набор значений (1-имеет, 0-не имеет) |
| updated | datetime |  |  |  | Время последнего обновления |
| deleted | bool |  |  |  | Строка удалена (перемещена в корзину) |

[Подробнее о запросах](https://developers.business.ru/api-polnoe/zaprosy_k_api/369)[Примеры запросов](https://developers.business.ru/api-polnoe/primery_zaprosov_k_api/377)[О дополнительных параметрах](https://developers.business.ru/api-polnoe/poluchenie_dannyh_iz_svyazannyh_modelej/454)Сформировать запрос к модели

Значения атрибутов (attributesforgoodsvalues)

**Описание модели:** Список значений атрибутов товаров и услуг

**URL ресурса:** https://myaccount.business.ru/api/rest/attributesforgoodsvalues.json

**Разрешенные запросы:** get(чтение), post(создание), put(изменение), delete(удаление)

| Параметр | Тип | Обязателен в запросах | Внешний ключ к модели | Значение по умолчанию | Описание |
| --- | --- | --- | --- | --- | --- |
| id | int | put, delete |  |  | Идентификатор |
| attribute\_id | int | post | [attributesforgoods](https://developers.business.ru/api-polnoe/atributy_attributesforgoods/294) |  | Ссылка на атрибут |
| name | string | post |  |  | Значение |
| updated | datetime |  |  |  | Время последнего обновления |

[Подробнее о запросах](https://developers.business.ru/api-polnoe/zaprosy_k_api/369)[Примеры запросов](https://developers.business.ru/api-polnoe/primery_zaprosov_k_api/377)[О дополнительных параметрах](https://developers.business.ru/api-polnoe/poluchenie_dannyh_iz_svyazannyh_modelej/454)Сформировать запрос к модели

Привязка атрибутов (goodsattributes)

**Описание модели:** Привязка атрибутов к товарам

**URL ресурса:** https://myaccount.business.ru/api/rest/goodsattributes.json

**Разрешенные запросы:** get(чтение), post(создание), put(изменение), delete(удаление)

| Параметр | Тип | Обязателен в запросах | Внешний ключ к модели | Значение по умолчанию | Описание |
| --- | --- | --- | --- | --- | --- |
| id | int | put, delete |  |  | Идентификатор |
| good\_id | int | post | [goods](https://developers.business.ru/api-polnoe/spravochnik_tovarov_i_uslug_goods/289) |  | Ссылка на товар |
| attribute\_id | int | post | [attributesforgoods](https://developers.business.ru/api-polnoe/atributy_attributesforgoods/294) |  | Ссылка на атрибут |
| value\_id | int |  | [attributesforgoodsvalues](https://developers.business.ru/api-polnoe/znacheniya_atributov_attributesforgoodsvalues/295) |  | Ссылка на значение атрибута |
| value | string |  |  |  | Значение атрибута |
| updated | datetime |  |  |  | Время последнего обновления |

[Подробнее о запросах](https://developers.business.ru/api-polnoe/zaprosy_k_api/369)[Примеры запросов](https://developers.business.ru/api-polnoe/primery_zaprosov_k_api/377)[О дополнительных параметрах](https://developers.business.ru/api-polnoe/poluchenie_dannyh_iz_svyazannyh_modelej/454)Сформировать запрос к модели

Модификации товаров

Данный подраздел содержит модели данных, обеспечивающие внешним приложениям доступ к информации о модификациях товаров и услуг.

Атрибуты модификаций (attributesformodifications)

**Описание модели:** Список атрибутов модификаций товаров и услуг

**URL ресурса:** https://myaccount.business.ru/api/rest/attributesformodifications.json

**Разрешенные запросы:** get(чтение), post(создание), put(изменение), delete(удаление)

| Параметр | Тип | Обязателен в запросах | Внешний ключ к модели | Значение по умолчанию | Описание |
| --- | --- | --- | --- | --- | --- |
| id | int | put, delete |  |  | Идентификатор |
| name | string | post |  |  | Наименование атрибута |
| selectable | bool | post |  |  | Флаг, определяющий имеет ли атрибут предопределенный набор значений (1-имеет, 0-не имеет) |
| updated | datetime |  |  |  | Время последнего обновления |
| deleted | bool |  |  |  | Строка удалена (перемещена в корзину) |

[Подробнее о запросах](https://developers.business.ru/api-polnoe/zaprosy_k_api/369)[Примеры запросов](https://developers.business.ru/api-polnoe/primery_zaprosov_k_api/377)[О дополнительных параметрах](https://developers.business.ru/api-polnoe/poluchenie_dannyh_iz_svyazannyh_modelej/454)Сформировать запрос к модели

Значения атрибутов (attributesformodificationsvalues)

**Описание модели:** Список значений атрибутов модификаций товаров и услуг

**URL ресурса:** https://myaccount.business.ru/api/rest/attributesformodificationsvalues.json

**Разрешенные запросы:** get(чтение), post(создание), put(изменение), delete(удаление)

| Параметр | Тип | Обязателен в запросах | Внешний ключ к модели | Значение по умолчанию | Описание |
| --- | --- | --- | --- | --- | --- |
| id | int | put, delete |  |  | Идентификатор |
| attribute\_id | int | post | [attributesformodifications](https://developers.business.ru/api-polnoe/atributy_modifikatsij_attributesformodifications/298) |  | Ссылка на атрибут |
| name | string | post |  |  | Значение |
| updated | datetime |  |  |  | Время последнего обновления |

[Подробнее о запросах](https://developers.business.ru/api-polnoe/zaprosy_k_api/369)[Примеры запросов](https://developers.business.ru/api-polnoe/primery_zaprosov_k_api/377)[О дополнительных параметрах](https://developers.business.ru/api-polnoe/poluchenie_dannyh_iz_svyazannyh_modelej/454)Сформировать запрос к модели

Привязка атрибутов (modificationsattributes)

**Описание модели:** Привязка атрибутов к модификациям товаров

**URL ресурса:** https://myaccount.business.ru/api/rest/modificationsattributes.json

**Разрешенные запросы:** get(чтение), post(создание), put(изменение), delete(удаление)

| Параметр | Тип | Обязателен в запросах | Внешний ключ к модели | Значение по умолчанию | Описание |
| --- | --- | --- | --- | --- | --- |
| id | int | put, delete |  |  | Идентификатор |
| goods\_modification\_id | int | post | [goodsmodifications](https://developers.business.ru/api-polnoe/privyazka_modifikatsij_goodsmodifications/301) |  | Ссылка на модификацию товара |
| attribute\_id | int | post | [attributesformodifications](https://developers.business.ru/api-polnoe/atributy_modifikatsij_attributesformodifications/298) |  | Ссылка на атрибут |
| value\_id | int |  | [attributesformodificationsvalues](https://developers.business.ru/api-polnoe/znacheniya_atributov_attributesformodificationsvalues/299) |  | Ссылка на значение атрибута |
| value | string |  |  |  | Значение атрибута |
| updated | datetime |  |  |  | Время последнего обновления |

[Подробнее о запросах](https://developers.business.ru/api-polnoe/zaprosy_k_api/369)[Примеры запросов](https://developers.business.ru/api-polnoe/primery_zaprosov_k_api/377)[О дополнительных параметрах](https://developers.business.ru/api-polnoe/poluchenie_dannyh_iz_svyazannyh_modelej/454)Сформировать запрос к модели

Привязка модификаций (goodsmodifications)

**Описание модели:** Привязка модификаций к товарам

**URL ресурса:** https://myaccount.business.ru/api/rest/goodsmodifications.json

**Разрешенные запросы:** get(чтение), post(создание), put(изменение), delete(удаление)

| Параметр | Тип | Обязателен в запросах | Внешний ключ к модели | Значение по умолчанию | Описание |
| --- | --- | --- | --- | --- | --- |
| id | int | put, delete |  |  | Идентификатор |
| name | string | post |  |  | Наименование модификации |
| good\_id | int | post | [goods](https://developers.business.ru/api-polnoe/spravochnik_tovarov_i_uslug_goods/289) |  | Ссылка на товар |
| part | string |  |  |  | Артикул |
| archive | bool |  |  |  | Флаг (0-товар не перемещен в архив, 1-товар перемещен в архив) |
| images | images |  |  |  | Изображения товара |
| remains\_min | float |  |  |  | Неснижаемый остаток на складе |
| updated | datetime |  |  |  | Время последнего обновления |
| deleted | bool |  |  |  | Строка удалена (перемещена в корзину) |

[Подробнее о запросах](https://developers.business.ru/api-polnoe/zaprosy_k_api/369)[Примеры запросов](https://developers.business.ru/api-polnoe/primery_zaprosov_k_api/377)[О дополнительных параметрах](https://developers.business.ru/api-polnoe/poluchenie_dannyh_iz_svyazannyh_modelej/454)Сформировать запрос к модели

Аналоги товаров (goodsanalogs)

**Описание модели:** Аналоги товаров

**URL ресурса:** https://myaccount.business.ru/api/rest/goodsanalogs.json

**Разрешенные запросы:** get(чтение), post(создание), put(изменение), delete(удаление)

| Параметр | Тип | Обязателен в запросах | Внешний ключ к модели | Значение по умолчанию | Описание |
| --- | --- | --- | --- | --- | --- |
| id | int | put, delete |  |  | Идентификатор |
| good\_id | int | post | [goods](https://developers.business.ru/api-polnoe/spravochnik_tovarov_i_uslug_goods/289) |  | Ссылка на товар |
| analog\_id | int | post | [goods](https://developers.business.ru/api-polnoe/spravochnik_tovarov_i_uslug_goods/289) |  | Ссылка на аналогичный товар |

[Подробнее о запросах](https://developers.business.ru/api-polnoe/zaprosy_k_api/369)[Примеры запросов](https://developers.business.ru/api-polnoe/primery_zaprosov_k_api/377)[О дополнительных параметрах](https://developers.business.ru/api-polnoe/poluchenie_dannyh_iz_svyazannyh_modelej/454)Сформировать запрос к модели

Комплектующие (kitcomponents)

**Описание модели:** Комплектующие в составе комлекта товаров

**URL ресурса:** https://myaccount.business.ru/api/rest/kitcomponents.json

**Разрешенные запросы:** get(чтение), post(создание), put(изменение), delete(удаление)

| Параметр | Тип | Обязателен в запросах | Внешний ключ к модели | Значение по умолчанию | Описание |
| --- | --- | --- | --- | --- | --- |
| id | int | put, delete |  |  | Идентификатор |
| kit\_id | int | post | [goods](https://developers.business.ru/api-polnoe/spravochnik_tovarov_i_uslug_goods/289) |  | Ссылка на комплект |
| good\_id | int | post | [goods](https://developers.business.ru/api-polnoe/spravochnik_tovarov_i_uslug_goods/289) |  | Ссылка на товар |
| modification\_id | int |  | [goodsmodifications](https://developers.business.ru/api-polnoe/privyazka_modifikatsij_goodsmodifications/301) |  | Ссылка на модификацию товара |
| measure\_id | int |  | [measures](https://developers.business.ru/api-polnoe/edinitsy_izmereniya_measures/349) |  | Ссылка на единицу измерения |
| amount | float | post |  |  | Количество |
| description | string |  |  |  | Описание |
| updated | datetime |  |  |  | Время последнего обновления |

[Подробнее о запросах](https://developers.business.ru/api-polnoe/zaprosy_k_api/369)[Примеры запросов](https://developers.business.ru/api-polnoe/primery_zaprosov_k_api/377)[О дополнительных параметрах](https://developers.business.ru/api-polnoe/poluchenie_dannyh_iz_svyazannyh_modelej/454)Сформировать запрос к модели

Поиск товаров (goodssearch)

**Описание модели:** Специальная модель для поиска товаров

**URL ресурса:** https://myaccount.business.ru/api/rest/goodssearch.json

**Разрешенные запросы:** get(чтение)

Модель goodssearch обеспечивает возможность релевантного поиска товаров (услуг, комплектов) по заданной строке. Модель повторяет алгоритм, реализованный в строках поиска товаров, которые используются во вкладках Товары / Услуги в карточках документов пользовательского интерфейса Бизнес.Ру. Модель возвращает массив, каждый элемент которого, представляет собой пару значений **идентификатор товара** - **идентификатор модификации товара** - всего не более 30 элементов.

Пример запроса:

$response = $api->request( “get”, ”goodssearch”, \[

"text" => "Искомая строка"

\] );

Сопутствующие товары (goodsrelated)
