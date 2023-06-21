

<template>
  <!-- สนใจ chartData เอาข้อมูลมาแสดง-->
  <Bar
    :options="chartOptions"
    :data="chartData"
  />

  <p>Check value from chartdata</p>
  <p>{{ this.chartData.datasets[0].data  }}</p>
  <p>{{ this.chartData.labels  }}</p>
  
</template>


<script>
// อ้างอิง
// https://medium.com/hackernoon/build-a-realtime-chart-with-vue-js-6527ac55c315
import { Bar } from 'vue-chartjs'
import { Chart as ChartJS, Title, Tooltip, Legend, BarElement, CategoryScale, LinearScale } from 'chart.js'
import { createApp } from 'vue'
import axios from 'axios'
ChartJS.register(Title, Tooltip, Legend, BarElement, CategoryScale, LinearScale)


export default{
  name: 'BarChart',
  components: { Bar },
  data () {
      
    return {
      // กำหนดข้อมูลเริ่มต้น ให้ chartData เป็น object ที่มี labels และ datasets ที่เป็น array ว่าง 
      chartData: {
        labels: [],
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
  mounted () {
    // เรียกใช้ method fillData เมื่อ component ถูก mount เสร็จ
    this.fillData()
  },

  methods: {
    // ดึงข้อมูลจาก api มาเก็บไว้ใน chartData
    fillData (){
      const get_data =  axios.get('http://localhost:8080/data').then(response => {
      // ข้อมูลจะอยู่ใน response.data
  //       {
          // "data": [
                // 10,
                // 20,
                // 30,
                // 40,
                // 50
                // ],
                // "labels": [
                // "January",
                // "February",
                // "March",
                // "April",
                // "May"
                // ]
                // }
      console.log(response.data.data)
      this.chartData = {
        labels: response.data.labels,
        datasets: [
          {
            label: 'Data One',
            backgroundColor: '#f87979',
            data: response.data.data
          }
        ]
      }
      console.log(this.chartData)
       
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
