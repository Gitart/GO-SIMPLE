
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
