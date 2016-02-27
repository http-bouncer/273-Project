/**
 * Created by tugceakin on 2/19/16.
 */

$( document ).ready(function() {
    console.log("script");
    $("#header").load("../views/header.html");
    //$("#navigation").load("/navigation.html");

    $('div.hidden').fadeIn(800).removeClass('hidden');

    $("[name='my-checkbox']").bootstrapSwitch({onColor: 'success', offColor: 'danger'});

    setTable();
    
    $( "#page1-form" ).submit(function( event ) {
	    	console.log("clicked submit button");

    	  // Stop form from submitting normally
    	  event.preventDefault();
    	 
    	  // Get some values from elements on the page:
    	  var $form = $( this );
    	  var term1 = $('#item1').val();
    	  var term2 = $('#item2').val();
    	 
    	  // Send the data using post
    	  var posting = $.post( "/postForm", { item1: term1, item2: term2 } );
    	 
    	  // Put the results in a div
    	  posting.done(function( data ) {
    	    console.log("YAY");
    	  });
    	});
});


function setTable(){
    var canvas = document.getElementById('chart1');
    if ($('#chart1').length != 0) {
        var ctx_co = canvas.getContext('2d'),
            startingData = {
                labels: [1, 2, 3, 4, 5, 6, 7],
                datasets: [
                    {
                        fillColor: "rgba(151,187,205,0.2)",
                        strokeColor: "rgba(151,187,205,1)",
                        pointColor: "rgba(151,187,205,1)",
                        pointStrokeColor: "#fff",
                        data: [0, 5, 10, 0, 0, 0, 0]
                    }
                ]
            },
            latestLabel = startingData.labels[6];

        var myChart = new Chart(ctx_co).Line(startingData, {animationSteps: 15});

        setInterval(function () {
            // Add two random numbers for each dataset
            console.log(Math.floor(Math.random() * 99) + 1);

            myChart.addData([Math.floor(Math.random() * 50) + 40], '');

            // Remove the first point so we dont just add values forever
            myChart.removeData();
        }, 3000);
    }
}
