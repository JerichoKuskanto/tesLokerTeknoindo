<template>
  <div class="table-container">
    <header class=" text-white p-4 flex justify-between items-center shadow-md gap-4">
      <router-link to="/insert">
        <button 
          class="bg-black hover:bg-gray-800 transition-all duration-300 text-white px-6 py-2 rounded shadow-lg border border-gray-700 
          hover:border-gray-500 focus:outline-none focus:ring-2 focus:ring-gray-600">
          Insert Transaction
        </button>
      </router-link>
      <input 
        v-model="searchQuery" 
        type="text" 
        placeholder="name or transaction number" 
        class="bg-black text-white placeholder-gray-500 px-4 py-2 rounded border border-gray-700 focus:outline-none focus:ring-2 focus:ring-gray-600" />
    </header>
    <div class="table-wrapper">
      <table class="styled-table">
        <thead>
          <tr>
            <th class="tightest">No</th>
            <th class="tighter">No Transaksi</th>
            <th class="tight">Tanggal</th>
            <th class="tight">Nama Customer</th>
            <th class="tighter">Jumlah Barang</th>
            <th class="tight">Sub Total</th>
            <th class="tight">Diskon</th>
            <th class="tight">Ongkir</th>
            <th class="tightest">Total</th>
          </tr>
        </thead>
        <tbody>
          <tr v-for="(transaksi, index) in filteredTransaksi" :key="index">
            <td>{{ index + 1 }}</td>
            <td>{{ transaksi.kode }}</td>
            <td>{{ transaksi.tgl }}</td>
            <td>
              <router-link 
                :to="{ name: 'editPage', params: { id: transaksi.salesID } }" 
                class="text-blue-500 hover:underline"
              >
                {{ transaksi.name }}
              </router-link>
            </td>

            <td>{{ transaksi.totalQty }}</td>
            <td>{{ formatRupiah(transaksi.subtotal) }}</td>
            <td>{{ formatRupiah(transaksi.diskon) }}</td>
            <td>{{ formatRupiah(transaksi.ongkir) }}</td>
            <td>{{ formatRupiah(transaksi.total_bayar) }}</td>
          </tr>
          <!-- Grand Total Row -->
          <tr>
            <td colspan="8" class="grand-total-label">Grand Total</td>
            <td class="grand-total-value">{{ formatRupiah(grandTotal) }}</td>
          </tr>
        </tbody>
      </table>
    </div>
  </div>
</template>

<script setup>
import { ref, computed, onMounted } from 'vue';
const transactionsData = ref([]);

onMounted(getTransactionData);
const apiUrl = process.env.VUE_APP_API_URL;

if (!apiUrl) {
  throw new Error("VUE_APP_API_URL is not set in the environment variables!");
}

async function getTransactionData() {
  const transactions = ref([]);

  

  var query = apiUrl + '/sales';
  const response = await fetch(query, {
      method: "GET",
      headers: {
          "Content-Type": "application/x-www-form-urlencoded",
      }
  });

  if (response.ok) {
      const data = await response.json();
      if (data.status == '200') {
          for (const key in data.data) {
              let sale = data.data[key];
              sale.tgl = dateTimeFormatter(sale.tgl);
              let custQuery = apiUrl+`/customer/${sale.custID}`;
              const custResponse = await fetch(custQuery, {
                  method: "GET",
                  headers: {
                      "Content-Type": "application/x-www-form-urlencoded",
                  }
              });

              if (custResponse.ok) {
                  const custData = await custResponse.json();
                  if (custData.status == '200') {
                      sale.name = custData.data.nama;
                  }
              }

              let salesDetQuery = apiUrl+`/salesDetSpec/${sale.salesID}`;
              const salesDetResponse = await fetch(salesDetQuery, {
                  method: "GET",
                  headers: {
                      "Content-Type": "application/x-www-form-urlencoded",
                  }
              });

              if (salesDetResponse.ok) {
                  const salesDetData = await salesDetResponse.json();
                  if (salesDetData.status == '200' && Array.isArray(salesDetData.data)) {
                      sale.totalQty = salesDetData.data.reduce((sum, item) => sum + (item.qty || 0), 0);
                  } else {
                      sale.totalQty = 0;
                  }
              } else {
                  sale.totalQty = 0;
              }
              transactions.value.push(sale);
          }

          transactionsData.value = transactions.value;
      }
  }
}

function dateTimeFormatter(timestamp) {
  const date = new Date(timestamp);
  return date.toLocaleDateString('en-GB', {
      day: '2-digit',
      month: 'long',
      year: 'numeric'
  });
}

const searchQuery = ref("");
const filteredTransaksi = computed(() => {
return transactionsData.value.filter((transaksi) =>
  transaksi.name.toLowerCase().includes(searchQuery.value.toLowerCase()) ||
  transaksi.kode.toLowerCase().includes(searchQuery.value.toLowerCase())
);
});

const grandTotal = computed(() => {
  return filteredTransaksi.value.reduce((sum, transaksi) => sum + (transaksi.total_bayar || 0), 0);
});

const formatRupiah = (value) => {
  if (!value) return "Rp. 0";
  return new Intl.NumberFormat("id-ID", {
    style: "currency",
    currency: "IDR",
    minimumFractionDigits: 0
  }).format(value);
};
</script>

<style scoped>
header {
  display: flex;
  align-items: center;
  gap: 1rem; 
}

body {
  background-color: #121212;
  color: #ffffff;
  font-family: Arial, sans-serif;
  margin: 0;
  padding: 0;
  overflow-x: hidden; 
}

.table-container {
  width: 100%;
  padding: 20px;
  display: flex;
  flex-direction: column;
  align-items: center;
}

.search-box {
  margin-bottom: 10px;
  padding: 8px;
  width: 100%; 
  max-width: 500px; 
  border: 1px solid #444;
  border-radius: 5px;
  background-color: #333;
  color: white;
}

.table-wrapper {
  width: 100%;
  max-width: 1200px;
  overflow-x: auto; 
  margin: 0 auto; 
  padding: 0;
  box-sizing: border-box;
}

.styled-table {
  width: 100%; 
  border-collapse: collapse;
  font-size: 0.9em;
  background-color: #1e1e1e; 
  box-shadow: 0px 4px 10px rgba(255, 255, 255, 0.1);
  table-layout: auto;
}

.styled-table thead {
  background-color: #333; 
  color: white; 
}

.styled-table th,
.styled-table td {
  padding: 10px 15px;
  text-align: left;
  border-bottom: 1px solid #444;
}

.styled-table .tightest {
  width: 5%;
}

.styled-table .tighter {
  width: 12%;
}

.styled-table .tight {
  width: 15%;
}

.styled-table tbody tr:nth-child(even) {
  background-color: #252525;
}

.styled-table tbody tr:hover {
  background-color: #333; 
}

.styled-table tbody td {
  color: #61dafb; 
}

.grand-total-label {
  text-align: right !important; 
  font-weight: bold;
  background-color: #252525;
  color: white;
  padding-right: 15px; 
}

.grand-total-value {
  font-weight: bold;
  background-color: #252525;
  color: #61dafb;
  text-align: right; 
}

@media (max-width: 768px) {
  .search-box {
    width: 90%; 
  }

  .table-wrapper {
    overflow-x: auto; 
    padding: 0; 
  }

  .styled-table {
    font-size: 0.8em;
    min-width: 600px; 
  }

  .styled-table th,
  .styled-table td {
    padding: 8px 10px; 
  }

  .grand-total-label {
    padding-right: 10px;
  }
}
</style>
