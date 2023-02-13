<script>
    import Home from "./Home.svelte";
    import dayjs from "dayjs";
    
    export let table_header_font = "";
	export let table_body_font = "";
    
    let token = localStorage.getItem("token");
    let akses_page = false;
    let listHome = [];
    let record = "";
    let record_message = "";
    let totalrecord = 0;

    async function initapp() {
        const res = await fetch("/api/valid", {
            method: "POST",
            headers: {
                "Content-Type": "application/json",
                Authorization: "Bearer " + token,
            },
            body: JSON.stringify({
                page: "DOMAIN-VIEW",
            }),
        });
        const json = await res.json();
        if (json.status === 400) {
            logout();
        } else if (json.status == 403) {
            alert(json.message);
        } else {
            akses_page = true;
            initHome();
        }
    }
    async function initHome() {
        const res = await fetch("/api/event", {
            method: "POST",
            headers: {
                "Content-Type": "application/json",
                Authorization: "Bearer " + token,
            },
            body: JSON.stringify({
            }),
        });
        const json = await res.json();
        if (json.status == 200) {
            record = json.record;
            record_message = json.message;
            if (record != null) {
                totalrecord = record.length;
                let no = 0
                let duration = 0;
                let duration_css = "";
                let status_css = "";
                for (var i = 0; i < record.length; i++) {
                    no = no + 1;
                    duration = diffdate(dayjs().format("YYYY-MM-DD HH:MM"),dayjs(record[i]["event_endevent"]).format("YYYY-MM-DD HH:MM"));
                    if(duration < 0){
                        duration = 0;
                        duration_css = "color:red;font-weight:bold;";
                    }else{
                        duration_css = "color:biru;font-weight:bold;";
                    }
                    if(record[i]["event_status"] == "ONLINE"){
                        status_css = "background:#FFEB3B;font-weight:bold;color:black;"
                    }else{
                        status_css = "background:#E91E63;font-weight:bold;color:white;"
                    }
                    listHome = [
                        ...listHome,
                        {
                            home_no: no,
                            home_id: record[i]["event_id"],
                            home_idwebsite: record[i]["event_idwebagen"],
                            home_websiteagen: record[i]["event_nmwebagen"],
                            home_name: record[i]["event_name"],
                            home_start: dayjs(record[i]["event_startevent"]).format("YYYY-MM-DD HH:MM"),
                            home_end: dayjs(record[i]["event_endevent"]).format("YYYY-MM-DD HH:MM"),
                            home_duration: duration,
                            home_durationcss: duration_css,
                            home_mindeposit: record[i]["event_mindeposit"],
                            home_money_in: record[i]["event_money_in"],
                            home_status: record[i]["event_status"],
                            home_status_css: status_css,
                            home_create: record[i]["event_create"],
                            home_update: record[i]["event_update"],
                        },
                    ];
                }
            }
        } else {
            logout();
        }
    }
    async function logout() {
        localStorage.clear();
        window.location.href = "/";
    }
    const handleRefreshData = (e) => {
        listHome = [];
        totalrecord = 0;
        setTimeout(function () {
            initHome();
        }, 500);
    };
    function diffdate(data1,data2){
        // const date1 = dayjs(data1).format("YYYY-MM-DD HH:MM")
        // const date2 = dayjs(data2).format("YYYY-MM-DD HH:MM")
        // let temp =date1.diff(date2,'day',true);
        // return temp;
        const date1 = dayjs(data2)
        const date2 = dayjs(data1)
        let temp = date1.diff(date2,'day',true);
      
        return parseInt(temp);
    }
    initapp()
</script>

{#if akses_page == true}
<Home
    on:handleRefreshData={handleRefreshData}
    {token}
    {table_header_font}
    {table_body_font}
    {listHome}
    {totalrecord}/>
{/if}