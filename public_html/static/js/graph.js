'use strict';

var optionslevels = {
																
    chart: {
        type: 'area',
        zoomType: 'x',

    },

    title: {
    text: 'Уровень воды'
    },

    plotOptions: {
    area: {

    marker: {
    radius: 1
    },
    lineWidth: 1,
    states: {
    hover: {
    lineWidth: 1
    }
    },
    threshold: null
    }
    },							


    xAxis: {
        type: 'datetime',
        
    },
    yAxis: {
        title: {
            text: 'Уровень, см'
        },
    },

    exporting: {
        enabled: false
    },   

    legend: {
        enabled: false
    },
    tooltip: {
        xDateFormat: '%Y-%m-%d %H:%M',
    }


}


function graphLevels(id) {
	

    var chartlevels = Highcharts.chart(id, optionslevels);	
    $.getJSON(
        '/api/levels/',
    function (data) {		
        let datas= [];		
        $.each(data.out,function(key,val) {									
            //console.log(val.Level)
            datas.push([val.Unix,val.Level]);
        });
        //console.log(datas);
        chartlevels.addSeries({
                id: "levelsID",
                name: "levelsName",
                data: datas
            
        });
        chartlevels.reflow();
                                                                                            
    });
}