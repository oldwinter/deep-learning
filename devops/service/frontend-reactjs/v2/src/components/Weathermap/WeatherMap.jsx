import React, { Component } from 'react';
import ReactDOM from 'react-dom';


import ReactEcharts from 'echarts-for-react';
// require("echarts/echarts.all")
//    require('../../3rd-party-lib/echarts.min')
//   var dom = document.getElementById("map");
//   var myChart = echarts.init(dom);


class WeatherMap extends Component {
    constructor(props) {
        super(props)
        
    }
    getOption=() => {
        require('echarts/map/js/china.js');

        let option = null
        var geoCoordMap = {
            "Beijing": [116.46, 39.92],
            "Hangzhou": [120.19, 30.26],
            "Shanghai": [121.48, 31.22],
        };

        var convertData = function (data) {
            var res = [];
            for (var i = 0; i < data.length; i++) {
                var geoCoord = geoCoordMap[data[i].name];
                if (geoCoord) {
                    res.push({
                        name: data[i].name,
                        value: geoCoord.concat(data[i].value)
                    });
                }
            }
            return res;
        };

        option = {
            // backgroundColor: '#404a59',
            title: {
                text: 'Air Quality',
                x: 'center',
                textStyle: {
                    color: '#fff'
                }
            },
            tooltip: {
                trigger: 'item',
                formatter: function (params) {
                    return params.name + ' : ' + params.value[2];
                }
            },
            legend: {
                orient: 'vertical',
                y: 'bottom',
                x: 'right',
                data: ['pm2.5'],
                textStyle: {
                    color: '#fff'
                }
            },
            visualMap: {
                min: 0,
                max: 200,
                calculable: true,
                inRange: {
                    color: ['#50a3ba', '#eac736', '#d94e5d']
                },
                textStyle: {
                    color: '#fff'
                }
            },
            geo: {
                map: 'china',
                label: {
                    emphasis: {
                        show: false
                    }
                },
                itemStyle: {

                    normal: {
                        opacity: 0.2,
                        areaColor: '#eee',
                        borderColor: '#111'
                    },
                    emphasis: {
                        areaColor: '#2a333d'
                    }
                }
            },
            series: [
                {
                    name: 'pm2.5',
                    type: 'scatter',
                    coordinateSystem: 'geo',
                    data: convertData([
                        { name: "Hangzhou", value: 24 },
                        { name: "Beijing", value: 179 },
                        { name: "Shanghai", value: 105 },
                    ]),
                    symbolSize: 12,
                    label: {
                        normal: {
                            show: false
                        },
                        emphasis: {
                            show: false
                        }
                    },
                    itemStyle: {
                        emphasis: {
                            borderColor: '#fff',
                            borderWidth: 1
                        }
                    }
                }
            ]
        }

        return option
    }

    render() {
        return (
            <div className="">
                {/* <div ref="pieChart" style={{width: "100%", height: "100%"}}></div> */}
                <ReactEcharts
                    option={this.getOption()}
                    style={{ height: '350px', width: '100%' }}
                    opts={{ renderer: 'svg' }} />
            </div>
        )

    }
}
export default WeatherMap;