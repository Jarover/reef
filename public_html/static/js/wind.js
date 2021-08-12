$(function () {

    let matas = new Map();
    let datas = []
    let start, interval;
   //console.log("!!!");

   $.getJSON('/api/chanels/3', function (data) {

        $.each(data.out,function(i,val) {		
            
            matas.set(val.Unix,[val.Level/100,0])
        });
        

    }).always(function() {
        $.getJSON('/api/chanels/2', function (data) {


            $.each(data.out,function(i,val) {	
                
                if (i == 0) {
                    start = val.Unix
                }
                if (i == 1) {
                    interval = val.Unix - start;
                    //console.log(interval);
                }

                d = matas.get(val.Unix);
                if (d) {
                    d[1] = val.Level/100
        
                    matas.set(val.Unix,d)
                    datas.push(d)
                }
            });
            //console.log(datas)


            chart = Highcharts.chart('wind', {

                title: {

                    text: 'Ветер', 
                  
                },
                legend: {
                    enabled: false
                },                
                yAxis: {
                    title: {
                        text: 'Сила ветра, м/c'
                    },
                },
                exporting: {
                    enabled: false
                },                
                xAxis: {

                    type: 'datetime',

                    offset: 40,

                },

                plotOptions: {

                    series: {

                        pointStart: start,

                        pointInterval: interval

                    }

                },

                series: [{

                    type: 'windbarb',

                    data: datas,

                    name: 'Wind',

                    
                    color: Highcharts.getOptions().colors[1],

                    showInLegend: false,

                    tooltip: {
                        xDateFormat: '%Y-%m-%d %H:%M',
                        valueSuffix: ' m/s'

                    }

                }, {

                    type: 'area',

                    keys: ['y', 'rotation'], // rotation is not used here

                    data: datas,

                    color: Highcharts.getOptions().colors[0],

                    fillColor: {

                        linearGradient: {

                            x1: 0,

                            x2: 0,

                            y1: 0,

                            y2: 1

                        },

                        stops: [

                            [0, Highcharts.getOptions().colors[0]],

                            [

                                1,

                                Highcharts.color(Highcharts.getOptions().colors[0])

                                .setOpacity(0.25).get()

                            ]

                        ]

                    },

                    name: 'Сила ветра',

                    tooltip: {
                        xDateFormat: '%Y-%m-%d %H:%M',
                        valueSuffix: ' m/s'

                    },

                    states: {

                        inactive: {

                            opacity: 1

                        }

                    }

                }]

            });
        });


    });


    });