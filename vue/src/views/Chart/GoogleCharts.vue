<template>
    <GChart
        type="ColumnChart"
        :data="chartData"
        :options="chartOptions"
    />
</template>

<script>
import {GChart} from 'vue-google-charts'
  
export default {
    extends: GChart,
    data(){
        return {            
            chartData: [
                ['Year', '입금', '출금'],
                ['2014', 1000, 400],
                ['2015', 1170, 460],
                ['2016', 660, 1120],
                ['2017', 1030, 540],
                ['2017', 1030, 540],
                ['2017', 1030, 540],
                ['2017', 1030, 540]
            ],
            chartOptions: {
            chart: {
            title: 'Company Performance',
            subtitle: 'Sales, Expenses, and Profit: 2014-2017',
            }
            }
        }
    },
    created(){
        this.chartData[1] = ['2014' , 200 , 500]
    }
    
}
</script>

<style>

</style>
