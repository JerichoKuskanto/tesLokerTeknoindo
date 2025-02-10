<template>
    <div class="p-4 text-white min-h-screen">
      <header class=" text-white p-4 flex justify-between items-center shadow-md gap-4">
        <router-link to="/">
          <button 
            class="bg-black hover:bg-gray transition-all duration-300 text-white px-6 py-2 rounded shadow-lg border border-gray-700 
            hover:border-gray-500 focus:outline-none focus:ring-2 focus:ring-gray-600">
            Back Home
          </button>
        </router-link>
        &nbsp;&nbsp;
        <button @click="deleteTransaction"
                class="bg-danger text-white px-3 py-2 rounded border border-dark">
            Delete Transaction
        </button>


      </header>

      <!-- Sales Info Table (Small Table) -->
      <div class="max-w-full overflow-x-auto mb-4">
        <table class="w-full border-collapse border border-gray-700 text-sm">
          <thead>
            <tr class="bg-gray-800 text-white">
              <th class="border border-gray-700 px-2 py-1">Tanggal</th>
              <th class="border border-gray-700 px-2 py-1">Customer</th>
              <th class="border border-gray-700 px-2 py-1">Nama</th>
              <th class="border border-gray-700 px-2 py-1">Telp</th>
            </tr>
          </thead>
          <tbody>
            <tr>
              <!-- Date Picker -->
              <td class="border border-gray-700 px-2 py-1">
                <input v-model="tanggal" type="date" class="border border-gray-700 px-2 py-1 w-full bg-black text-white">
              </td>
  
              <!-- Customer Dropdown -->
              <td class="border border-gray-700 px-2 py-1">
                <select v-model="selectedCustomer" @change="updateCustomerInfo" class="border border-gray-700 px-2 py-1 w-full bg-black text-white">
                  <option value="" disabled>Select Customer</option>
                  <option 
                    v-for="customer in customers" 
                    :key="customer.customerID" 
                    :value="customer.customerID">
                    {{ customer.kode }}
                  </option>
                </select>
              </td>
  
              <!-- Auto-filled Nama -->
              <td class="border border-gray-700 px-2 py-1">
                <input v-model="customerNama" class="border border-gray-700 px-2 py-1 w-full bg-black text-white" disabled>
              </td>
  
              <!-- Auto-filled Telp -->
              <td class="border border-gray-700 px-2 py-1">
                <input v-model="customerTelp" class="border border-gray-700 px-2 py-1 w-full bg-black text-white" disabled>
              </td>
  
            </tr>
          </tbody>
        </table>
      </div>
  
      <!-- Sales Details Table (Main Table) -->
      <div class="max-w-full overflow-x-auto">
        <table class="w-full border-collapse border border-gray-700 text-sm">
          <thead>
            <tr class="bg-gray-800 text-white">
              <th class="border border-gray-700 px-2 py-1">No</th> <!-- Added "No" Column -->
              <th class="border border-gray-700 px-2 py-1">Barang ID</th>
              <th class="border border-gray-700 px-2 py-1">Nama Barang</th>
              <th class="border border-gray-700 px-2 py-1">Harga Bandrol</th>
              <th class="border border-gray-700 px-2 py-1">Qty</th>
              <th class="border border-gray-700 px-2 py-1">Diskon %</th>
              <th class="border border-gray-700 px-2 py-1">Diskon Nilai</th>
              <th class="border border-gray-700 px-2 py-1">Harga Diskon</th>
              <th class="border border-gray-700 px-2 py-1">Total</th>
              <th class="border border-gray-700 px-2 py-1">Action</th>
            </tr>
          </thead>
  
          <tbody>
            <tr v-for="(row, index) in tableData" :key="index">
              <td class="border border-gray-700 px-2 py-1 text-center">
                {{ index + 1 }} <!-- Auto-increment No -->
              </td>
              <td class="border border-gray-700 px-2 py-1">
                <select v-model="row.barangID" @change="updateBarangInfo(row)" class="border border-gray-700 px-2 py-1 w-full bg-black text-white">
                  <option value="" disabled>Select Barang</option>
                  <option 
                    v-for="barang in barangs" 
                    :key="barang.barangID" 
                    :value="barang.barangID">
                    {{ barang.kode }}
                  </option>
                </select>
              </td>
              <td class="border border-gray-700 px-2 py-1">
                <input v-model="row.barangNama" class="border border-gray-700 px-2 py-1 w-full bg-black text-white" disabled>
              </td>
              <td class="border border-gray-700 px-2 py-1">
                {{ formatRupiah(row.hargaBandrol) }}
              </td>
              <td class="border border-gray-700 px-2 py-1">
                <input v-model="row.qty" type="number" class="border border-gray-700 px-2 py-1 w-full bg-black text-white">
              </td>
              <td class="border border-gray-700 px-2 py-1">
                <input v-model="row.diskonPct" @input="updateDiskon(row)" type="number" step="0.01" class="border border-gray-700 px-2 py-1 w-full bg-black text-white">
              </td>
              <td class="border border-gray-700 px-2 py-1">
                {{ formatRupiah(row.diskonNilai) }}
              </td>
              <td class="border border-gray-700 px-2 py-1">
                {{ formatRupiah(row.hargaDiskon) }}
              </td>
              <td class="border border-gray-700 px-2 py-1">
                {{ formatRupiah(row.total) }}
              </td>
              <td class="border border-gray-700 px-2 py-1 text-center">
                <button @click="removeRow(index)" class="bg-black text-white border border-gray-700 px-2 py-1 rounded text-xs">
                  X
                </button>
              </td>
            </tr>
          </tbody>
          <tfoot>
            <!-- Sub Total Row -->
            <tr>
              <td colspan="8" class="border border-gray-700 px-2 py-1 text-right grand-total-label">Sub Total</td>
              <td class="border border-gray-700 px-2 py-1">
                {{ formatRupiah(subTotal) }}
              </td>
              <td class="border border-gray-700 px-2 py-1"></td>
            </tr>
  
            <!-- Diskon Row -->
            <tr>
              <td colspan="8" class="border border-gray-700 px-2 py-1 text-right grand-total-label">Diskon</td>
              <td class="border border-gray-700 px-2 py-1">
                <input v-model.number="diskon" type="number" class="border border-gray-700 px-2 py-1 w-full bg-black text-white">
              </td>
              <td class="border border-gray-700 px-2 py-1"></td>
            </tr>
  
            <!-- Ongkir Row -->
            <tr>
              <td colspan="8" class="border border-gray-700 px-2 py-1 text-right grand-total-label">Ongkir</td>
              <td class="border border-gray-700 px-2 py-1">
                <input v-model.number="ongkir" type="number" class="border border-gray-700 px-2 py-1 w-full bg-black text-white">
              </td>
              <td class="border border-gray-700 px-2 py-1"></td>
            </tr>
  
            <!-- Total Bayar Row -->
            <tr>
              <td colspan="8" class="border border-gray-700 px-2 py-1 text-right font-bold grand-total-label">Total Bayar</td>
              <td class="border border-gray-700 px-2 py-1 font-bold">
                {{ formatRupiah(totalBayar) }}
              </td>
              <td class="border border-gray-700 px-2 py-1"></td>
            </tr>
          </tfoot>
  
  
        </table>
      </div>
  
      <!-- Add Row Button -->
      <div class="text-center mt-2">
        <button @click="addRow" class="bg-black text-white px-3 py-1 rounded text-sm border border-gray-700">
          + Add Row
        </button>
      </div>
  
      <!-- Submit Button -->
      <button @click="submitForm" class="bg-black text-white px-3 py-1 rounded mt-2 text-sm border border-gray-700">
        Submit
      </button>
    </div>
  </template>
  
  <script setup>
  import { ref, onMounted, computed } from 'vue';
  import { useRoute, useRouter } from 'vue-router'

  const apiUrl = import.meta.env.VITE_API_URL || "http://localhost:8181"; // Default to localhost for development

  const route = useRoute()
  const router = useRouter()
  const idEdit = route.params.id 

  const salesData = ref([]);
  const salesDetailData = ref([]);

  async function getAllTransaction() {
    var query = apiUrl + '/sales/' + idEdit;
    const response = await fetch(query, { method: "GET", headers: { "Content-Type": "application/x-www-form-urlencoded" } });
    if (response.ok) {
        const data = await response.json();
        if (data.status == '200') {
            salesData.value = data.data;

            var queryDetail = apiUrl +'/salesDetSpec/' + idEdit;
            const responseDetail = await fetch(queryDetail, { method: "GET", headers: { "Content-Type": "application/x-www-form-urlencoded" } });

            if (responseDetail.ok) {
                const data2 = await responseDetail.json();
                if (data2.status == '200') {
                  salesDetailData.value = data2.data || [];
                    tableData.value = salesDetailData.value.length > 0 
                        ? salesDetailData.value.map((item) => ({
                            barangID: item.barang_ID || '',
                            qty: item.qty || 1,
                            diskonPct: item.diskon_pct || 0,
                        }))
                        : [{ barangID: '', barangNama: '', hargaBandrol: '', qty: '', diskonPct: '', diskonNilai: '', hargaDiskon: '', total: '' }];

                    tableData.value.forEach(row => updateBarangInfo(row));
                    tableData.value.forEach(row => updateDiskon(row));
                } else {
                    console.error("Sales Detail Failed!", data2.message);
                }
            }
        } else {
            console.error("Sales Failed!", data.message);
        }
    }

    tanggal.value = formatDate(salesData.value.tgl);
    selectedCustomer.value = salesData.value.custID;
    diskon.value = salesData.value.diskon;
    ongkir.value = salesData.value.ongkir;

    updateCustomerInfo();
  }



  function formatDate(dateString) {
  return dateString.split("T")[0]; 
}
  
  const tanggal = ref("");
  const selectedCustomer = ref("");
  const customerNama = ref("");
  const customerTelp = ref("");
  const customers = ref([]);
  const barangs = ref([]);
  
  const tableData = ref([
    { barangID: '', barangNama: '', hargaBandrol: '', qty: '', diskonPct: '', diskonNilai: '', hargaDiskon: '', total: '' }
  ]);
  
  function addRow(){
    tableData.value.push({
      barangID: '', barangNama: '', hargaBandrol: '', qty: '', diskonPct: '', diskonNilai: '', hargaDiskon: '', total: ''
    });
  }

  const diskon = ref(0);
  const ongkir = ref(0);
  
  function updateDiskon(row){
    if (row.hargaBandrol && row.qty && row.diskonPct) {
      row.diskonNilai = (parseFloat(row.hargaBandrol) * parseInt(row.qty) * parseFloat(row.diskonPct)) / 100;
    } else {
      row.diskonNilai = 0;
    }
  
    if (row.hargaBandrol) {
      row.hargaDiskon = parseFloat(row.hargaBandrol) - (parseFloat(row.diskonNilai)/parseInt(row.qty));
    } else {
      row.hargaDiskon = 0;
    }
  
    if (row.qty && row.hargaDiskon) {
      row.total = parseInt(row.qty) * parseFloat(row.hargaBandrol) - parseFloat(row.diskonNilai);
    } else {
      row.total = 0;
    }
  }
  
  const subTotal = computed(() => {
    return tableData.value.reduce((sum, item) => sum + (parseFloat(item.total) || 0), 0);
  });
  
  const totalBayar = computed(() => {
    return subTotal.value - (parseFloat(diskon.value) || 0) + (parseFloat(ongkir.value) || 0);
  });
  
  function formatRupiah(value){
    if (!value) return "Rp. 0";
    return new Intl.NumberFormat("id-ID", {
      style: "currency",
      currency: "IDR",
      minimumFractionDigits: 0
    }).format(value);
  }
  
  async function deleteTransaction() {
    const confirmation = confirm("Are you sure you want to delete this transaction?");
    if (!confirmation) {
        console.log("Deletion canceled.");
        return;
    }

    console.log(`Deleting existing sales details for salesID: ${idEdit}`);
    const deleteDetailResponse = await fetch(apiUrl+ `/salesDet/` + idEdit, {
        method: "DELETE",
        headers: { "Content-Type": "application/x-www-form-urlencoded" },
    });

    if (deleteDetailResponse.ok) {
        const deleteDetail = await deleteDetailResponse.json();
        if (deleteDetail.status == '200') {
            const deleteSalesResponse = await fetch(apiUrl + `/sales/` + idEdit, {
                method: "DELETE",
                headers: { "Content-Type": "application/x-www-form-urlencoded" },
            });

            if (deleteSalesResponse.ok) {
                const deleteSales = await deleteSalesResponse.json();
                if (deleteSales.status == '200') {
                    console.log("Success Deleting ");
                } else {
                    console.error("Failed!", deleteSales.message);
                }
            }
        } else {
            console.error("Failed!", deleteDetail.message);
        }
    }
    router.push("/");
  }



  async function getAllCustomer() {
    var query = apiUrl+ '/customer';
    const response = await fetch(query, { method: "GET", headers: { "Content-Type": "application/x-www-form-urlencoded" } });
    if (response.ok) {
      const data = await response.json();
      if (data.status == '200') {
        customers.value = data.data;
      } else {
        console.error("Failed!", data.message);
      }
    }
  }
  
  async function getAllBarang() {
    var query = apiUrl + '/barang';
    const response = await fetch(query, { method: "GET", headers: { "Content-Type": "application/x-www-form-urlencoded" } });
    if (response.ok) {
      const data = await response.json();
      if (data.status == '200') {
        barangs.value = data.data;
      } else {
        console.error("Failed!", data.message);
      }
    }
  }
  
  function updateCustomerInfo(){
    const selected = customers.value.find(c => c.customerID == selectedCustomer.value);
    if (selected) {
      customerNama.value = selected.nama;
      customerTelp.value = selected.telp;
    } else {
      customerNama.value = "";
      customerTelp.value = "";
    }
  }
  
  function updateBarangInfo(row){
    const selected = barangs.value.find(b => b.barangID == row.barangID);
  
    if (selected) {
      row.barangNama = selected.nama;
      row.hargaBandrol = selected.harga;
    } else {
      row.barangNama = "";
      row.hargaBandrol = "";
    }
  }
  
  
  async function submitForm() {
    try {
        const formDataSales = new URLSearchParams({
            kode: salesData.value.kode,
            tgl: tanggal.value,
            custID: selectedCustomer.value,
            subtotal: subTotal.value,
            diskon: diskon.value || 0,
            ongkir: ongkir.value || 0,
            total: totalBayar.value,
        });

        const salesResponse = await fetch(apiUrl + `/sales/` + idEdit, {
            method: "PUT",
            headers: { "Content-Type": "application/x-www-form-urlencoded" },
            body: formDataSales,
        });

        const salesDataResponse = await salesResponse.json();
        console.log("Sales API Response:", salesDataResponse);

        console.log(`Deleting existing sales details for salesID: ${idEdit}`);
        const deleteResponse = await fetch(apiUrl +`/salesDet/`+ idEdit, {
            method: "DELETE",
            headers: { "Content-Type": "application/x-www-form-urlencoded" },
        });

        const deleteData = await deleteResponse.json();
        console.log("Delete Sales Details Response:", deleteData);

        for (const row of tableData.value) {
            const formDataSalesDetail = new URLSearchParams({
                salesID: idEdit,
                barangID: row.barangID,
                hargaBandrol: row.hargaBandrol,
                qty: row.qty,
                diskonPct: row.diskonPct,
                diskonNilai: row.diskonNilai,
                hargaDiskon: row.hargaDiskon,
                total: row.total,
            });

            const insertResponse = await fetch(apiUrl + "/salesDet", {
                method: "POST",
                headers: { "Content-Type": "application/x-www-form-urlencoded" },
                body: formDataSalesDetail,
            });

            const insertData = await insertResponse.json();
            console.log("Insert Sales Detail Response:", insertData);

        }
        alert("Insert successful!");
        router.push("/");
        } catch (error) {
            console.error("Error during update:", error);
        }
  }

  function removeRow(index){
    tableData.value.splice(index, 1);
  }
  
  onMounted(getAllCustomer);
  onMounted(getAllBarang);
  onMounted(getAllTransaction)
  </script>
  
  <style scoped>
  .grand-total-label {
    text-align: right !important; 
    font-weight: bold;
    background-color: #252525; 
    color: white;
    padding-right: 15px; 
  }
  
  .max-w-full {
    width: 100%;
    overflow-x: auto; 
    -webkit-overflow-scrolling: touch; 
  }
  
  .max-w-full table {
    width: 100%;
    min-width: 800px; 
    border-collapse: collapse;
  }
  
  .max-w-full th,
  .max-w-full td {
    white-space: nowrap; 
    padding: 8px;
  }
  
  table {
    border-collapse: collapse;
  }
  
  th,
  td {
    text-align: left;
    padding: 8px;
  }
  
  th {
    background-color: #252525;
    color: white;
  }
  </style>
  