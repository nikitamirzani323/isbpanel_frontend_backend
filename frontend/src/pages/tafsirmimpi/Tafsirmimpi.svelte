<script>
    import Home from "./Home.svelte";
   
    
   export let table_header_font = "";
   export let table_body_font = "";
   
    let token = localStorage.getItem("token");
    let akses_page = false;
    let listHome = [];
    let listPage = [];
    let search = "";
    let sData = "";
    let record = "";
    let record_message = "";
    let perpage = 0;
    let page = 0;
    let totalrecord = 0;
    let totalrecordall = 0;
    let totalpaging = 0;

   async function initapp() {
       const res = await fetch("/api/valid", {
           method: "POST",
           headers: {
               "Content-Type": "application/json",
               Authorization: "Bearer " + token,
           },
           body: JSON.stringify({
               page: "TAFSIRMIMPI-VIEW",
           }),
       });
       const json = await res.json();
       if (json.status === 400) {
           logout();
       } else if (json.status == 403) {
           alert(json.message);
       } else {
            akses_page = true;
           initHome("");
       }
   }
   async function initHome(e) {
        listHome = [];
       const res = await fetch("/api/tafsirmimpi", { 
           method: "POST",
           headers: {
               "Content-Type": "application/json",
               Authorization: "Bearer " + token,
           },
           body: JSON.stringify({
                tafsirmimpi_search: e,
                tafsirmimpi_page : parseInt(page)
           }),
       });
       const json = await res.json();
       if (json.status == 200) {
            record = json.record;
            record_message = json.message;
            perpage = json.perpage;
            totalrecordall = json.totalrecord;
            if (record != null) {
                totalpaging = Math.ceil(parseInt(totalrecordall) / parseInt(perpage))
                totalrecord = totalrecordall;
                let no = 0
                if(page > 1){
                    no = parseInt(page) 
                }
                for (var i = 0; i < record.length; i++) {
                   no = no + 1;
                   listHome = [
                       ...listHome,
                       {
                            tafsirmimpi_no: no,
                            tafsirmimpi_id: record[i]["tafsirmimpi_id"],
                            tafsirmimpi_mimpi: record[i]["tafsirmimpi_mimpi"],
                            tafsirmimpi_artimimpi: record[i]["tafsirmimpi_artimimpi"],
                            tafsirmimpi_angka2d: record[i]["tafsirmimpi_angka2d"],
                            tafsirmimpi_angka3d: record[i]["tafsirmimpi_angka3d"],
                            tafsirmimpi_angka4d: record[i]["tafsirmimpi_angka4d"],
                            tafsirmimpi_status: record[i]["tafsirmimpi_status"],
                            tafsirmimpi_statuscss: record[i]["tafsirmimpi_statuscss"],
                            tafsirmimpi_create: record[i]["tafsirmimpi_create"],
                            tafsirmimpi_update: record[i]["tafsirmimpi_update"],
                       },
                   ];
                }
                listPage = [];
                for(var i=1;i<totalpaging;i++){
                    listPage = [
                        ...listPage,
                        {
                            page_id: i,
                            page_value: ((i*perpage)-perpage),
                            page_display: i + " Of " + perpage*i,
                        },
                    ];
                }
           }
       } else {
           alert("Error")
       }
   }
   async function logout() {
       localStorage.clear();
       window.location.href = "/";
   }
   const handleTafsirMimpi = (e) => {
        search = e.detail.searchTafsirMimpi;
        initHome(search)
   };
   const handleRefreshData = (e) => {
       listHome = [];
       totalrecord = 0;
       setTimeout(function () {
           initHome();
       }, 500);
   };
   const handlePaging = (e) => {
        page = e.detail.page
        initHome("")
    };
   initapp()
</script>
{#if akses_page == true}
<Home
    on:handlePaging={handlePaging}
    on:handleTafsirMimpi={handleTafsirMimpi}
    on:handleRefreshData={handleRefreshData}
    {token}
    {table_header_font}
    {table_body_font}
    {listHome}
    {listPage}
    {totalrecord}/>
{/if}