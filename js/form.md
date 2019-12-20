## Load function

```js
$(document).ready(function() {
          RefreshTable();	     
          $('#Inter').datepicker({format: 'dd.mm.yyyy'});
          $('#Intersecond').datepicker({format: 'dd.mm.yyyy'});
});
```

## Refresh table 

```js     
function RefreshTable(){
    $('#rls').DataTable({
	 "language": {
	        "decimal":       ",",
		"thousands":     ".",
		"search":        "Пошук ",
		"lengthMenu":    "Display _MENU_ records per page",
                                        "zeroRecords":   "Даних не знайдено",
                                        "info":          "Стр _PAGE_ з _PAGES_",
                                        "infoEmpty":     "Немаэ доступник записів",
                                        "infoFiltered":  "(filtered from _MAX_ total records)",
			  	        "columnDefs":    [{"visible": false, "targets":2}]
			},
    				        "scrollY":       "600px",
                                        "scrollCollapse": true,
				                        "paging":         false,
				                        "ordering":       true,
				                        "info":           false,
				                        "order":          [[2, "asc" ]],

		});
     }
```

## Grouping in table

```js
 $(document).ready(function() {
    $('#rls').DataTable({
	"language": {
	"decimal":       ",",
	"thousands":     ".",
	"search":        "Поиск ",
	"lengthMenu":    "Display _MENU_ records per page",
        "zeroRecords":   "Даних не знайдено",
        "info":          "Стр _PAGE_ з _PAGES_",
        "infoEmpty":     "Немаэ доступник записів",
        "infoFiltered":  "(filtered from _MAX_ total records)",
	"columnDefs":    [{"visible": false, "targets":2}],
        "scrollY":       "650px"},
	"drawCallback":   function (settings) {
                var api  = this.api();
                var rows = api.rows({page:'current'}).nodes();
                var last = null;
                api.column(2, {page:'current'} ).data().each(function(group,i) {
                if (last !== group ) {
                      $(rows).eq(i).before('<tr class="group"><td colspan="7"> <i class="fas fa-calendar-alt"></i> ' + group + '</td></tr>');
                last   = group;
                }
                });
                },
                "scrollCollapse": true,
		        "paging":         false,
			"ordering":       true,
			"info":           false,
			"order":          [[2, "asc" ]],
			"dom":            '<"toolbar">frtip'
	 });


					    
	$('#rls tbody').on( 'click', 'tr.group', function () {
	    var currentOrder = table.order()[0];

	    if (currentOrder[0] === 2 && currentOrder[1] === 'asc' ) {
	        table.order([2,'desc']).draw();
		}
	    else {
	        table.order([2,'asc']).draw();
		}
   });

    $("div.toolbar").html('<h4> <i class="fas fa-calendar-alt"></i> Техн специалисти</h4><b>Техничні спеціалисти Дата:2019-12-20 19:07:41</b>');
});
```					


## Edit form

```js
    function Form_edit(Id){
      	$("#form_candidat").show(100);
      	$("#number").text(Id);
         urlid="/hrm/candidatedit/"+Id;
         

          $.ajax({type:'GET',dataType: 'json', url: urlid, success: function(response) {
                $('#Fam').val(response.Fam);
                $('#Name').val(response.Name);
                $('#Vacation').val(response.Vacation);
                $('#Staj').val(response.Staj);
                $('#Mob').val(response.Mob);
                $('#Note').val(response.Note);
          }});
    }
```

## Hide form

```js
function Form_hide(){
  	window.location.reload();
  	$("#form_candidat").hide(100);
}
```

## Save form data to service
    
```js
function Savecandidatform(){
   var str = $("#Candidatsform").serialize();
   $.ajax({type:'POST', url: '/hrm/candidatinput/', data:str, success: function(response) {
                $('#Fam').focus();
                $('#Candidatsform')[0].reset();
                $.notify(response, "success");
          }});
    }
```

## Delete function

```js                
function Deletet(Id){
	
	var ur='/hrm/candidatdelete/'+Id;
	$.ajax({type:'POST', url: ur,  success: function(response) {
                $.notify(response, "success");
                window.location.reload();
                 

          }});
}
```


## Current date
```js
script type="text/javascript">
   var time = new Date();
   var year = time.getFullYear();
   $("#pyear").html("&copy " + year + " Unity-Bars");
</script>
```	 


## Buttons
```html
<div class="panel-footer clearfix">
   <div class="pull-right" style="text-align: right; padding: 20px;">
      <button type="reset"   class="btn btn-info"     id="Reject"   onclick="Form_hide();">Вийти</button>
      <button type="button"  class="btn btn-default"  id="Inform"   onclick="Information();">Довідка</button>
      <button type="button"  class="btn btn-success"  id="Save"     onclick="Savecandidatform()">Зберегти</button>
  </div>
</div>
```


