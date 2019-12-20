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


