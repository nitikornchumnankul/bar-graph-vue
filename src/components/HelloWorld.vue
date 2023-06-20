

<template>
  <Bar
    :options="chartOptions"
    :data="chartData"
    v-on:="this.load_data()"
  />
  <button @click="load_data">Load Data</button>


  <p>Check value from chartdata</p>
  <p>{{ this.chartData.datasets[0].data  }}</p>
  <p>{{ this.chartData.labels  }}</p>

  <p>Check value from chartdataNew</p>
<!-- 
  <p>{{ this.chartDataNew.datasets[0].data  }}</p>
  <p>{{ this.chartDataNew.labels  }}</p> -->

  
</template>


<script>
import { Bar } from 'vue-chartjs'
import { Chart as ChartJS, Title, Tooltip, Legend, BarElement, CategoryScale, LinearScale } from 'chart.js'
import { createApp } from 'vue'
import axios from 'axios'
ChartJS.register(Title, Tooltip, Legend, BarElement, CategoryScale, LinearScale)


export default{
  name: 'BarChart',
  components: { Bar },
  setup(){
    
    return {
      data: [],
      temp: [],
      chartData: {
        labels: ['January', 'February', 'March', 'April', 'May'],
        datasets: [
          {
            label: 'Data One',
            backgroundColor: '#f87979',
            data: []
          }
        ]
      },
    }
  },
  data () {
      
    return {
      chartDataNew: {
        labels: this.chartData.labels,
        datasets: [
          {
            label: 'Data One',
            backgroundColor: '#f87979',
            data: this.chartData.datasets[0].data
          }
        ]
      },
    }
  },
  methods: {
    load_data: function(){
      const get_data =  axios.get('http://localhost:8080/data').then(response => {
      this.temp = response.data.data
      this.chartData.datasets[0].data = response.data.data
      this.chartData.datasets[0].label = "Data One"
      this.chartData.labels = ['January', 'February', 'March', 'April', 'May']
      console.log(this.chartData.datasets[0].data)
      console.log(this.chartData.labels)
       
    })
    }
  },
}


</script>
<style scoped>
h1 {
  font-weight: 500;
  font-size: 2.6rem;
  position: relative;
  top: -10px;
}

h3 {
  font-size: 1.2rem;
}

.greetings h1,
.greetings h3 {
  text-align: center;
}

@media (min-width: 1024px) {
  .greetings h1,
  .greetings h3 {
    text-align: left;
  }
}
</style>
