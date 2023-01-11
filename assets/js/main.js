

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
    
    $('#csv_syuturyoku').submit(function() {
      var yoso_class_name = "";
      var Show_ClassName_ary = [];
      var T_class_name = "";
      var judge = false
      var yoso_id = "";
      
      //  プレビューエリアの子要素のループ
      $("#Preview_table").children().each(function(index, element){
         
         $(this).children().each(function(index, element){
            //  一時配列の作成
            var tmp_ary =[];
            //  行ループ
            $(this).children().each(function(index, element){
               //  ループ要素のid取得
               yoso_id = $(this).attr("id");
               yoso_class_name = $(this).attr("class");
               if ($('.' + yoso_class_name).css('display') != 'none') {
                  tmp_ary.push($('#' + yoso_id).text().replace(/\n/g, '').replace(/\s+/g,''));
               } 
             })
             
             Show_ClassName_ary.push(tmp_ary)
          })
        
       })
       
      const Form_csv = document.getElementById("csv_syuturyoku");
      
      var input_box = document.createElement("input");
      input_box.name = "Csv_Vals";
      input_box.type = "hidden";
      input_box.value = Show_ClassName_ary;
      Form_csv.appendChild(input_box);
      var input_box2 = document.createElement("input");
      input_box2.name = "Csv_Vals_cnt";
      input_box2.type = "hidden";
      input_box2.value = Show_ClassName_ary[0].length;
      Form_csv.appendChild(input_box2);
    
      // var input_box2 = document.createElement("input");
      // input_box2.name = "a";
      // input_box2.type = "hidden";
      // input_box2.value = "test";
      // Form_csv.appendChild(input_box2);
       
      //  Show_ClassName_ary.forEach(element => {
      //    var input_box = document.createElement("input");
      //    input_box.name = "Csv_Vals[]";
      //    input_box.type = "hidden";
      //    input_box.value = element;
      //    Form_csv.appendChild(input_box);
      //  });
     
      // // //  $("#csv_date").val(Show_ClassName_ary)
      // // //  $("#col_cnt").val(Show_ClassName_ary[0].length)
      // return false;
   
  });
});


function Clone_Yoso(copy_id) {
   
   if($('#Preview_table').length != 1){
      $("#Edit_Table").clone().appendTo('#Preview').attr("id","Preview_table");
      $(".row_0 th:last").remove();
      $("#Preview_table tr th,#Preview_table tr td").hide();
      
      var i = 1;
      $("#Preview_table tr").children().each(function(index, element){
        var class_name = $(this).attr("class").replace("ui-draggable ui-draggable-handle","");
        $(this).attr("class","ple_" + class_name);
        $(this).attr("id","ple_" + class_name.replace(" ","") + "_r" + i);
        i++;
      })
      
	}

   var el = document.getElementById('Preview_table');
   var dragger = tableDragger.default(el, {
         dragHandler: "." + "ple_" + copy_id + "_02"
   })
   
   $("." + "ple_" + copy_id+"_02").show();
  
}