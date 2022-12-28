
$(function(){
   $('th').draggable({
      helper: 'clone',
	});
   
   $("#Preview").droppable({
      //ドロップOKの要素を指定
      accept :"th" ,
      //ドロップ時の動作
      drop : function(event , ui){
         var draggable = ui.draggable;
         Clone_Yoso(draggable.attr("id"))
      }
    });
    
    $('#Preview_table th').draggable({
      start: function(){
        alert(9);
     }
	});  
});



function Clone_Yoso(copy_id) {
   
   if($('#Preview_table').length != 1){
      $("#Edit_Table").clone().appendTo('#Preview').attr("id","Preview_table");
      $(".row_0 th:last").remove();
      $("#Preview_table tr th,#Preview_table tr td").hide();
      $("#Preview_table tr").children().each(function(index, element){
        var class_name = $(this).attr("class").replace("ui-draggable ui-draggable-handle","");
        $(this).attr("class",class_name)
      })
	}

   var el = document.getElementById('Preview_table');
   var dragger = tableDragger.default(el, {
         dragHandler: "." + copy_id + "_02"
   })
   
   $("." + copy_id+"_02").show();
  
}