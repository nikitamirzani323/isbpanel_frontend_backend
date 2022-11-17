<script>
  import { Input } from "sveltestrap";
  import dayjs from "dayjs";
  import Panel from "../../components/Panel.svelte";
  import Loader from "../../components/Loader.svelte";
  import Button from "../../components/Button.svelte";
  import Modal from "../../components/Modal.svelte";
  import { createEventDispatcher } from "svelte";

  export let table_header_font = "";
  export let table_body_font = "";
  export let token = "";
  export let listHome = [];
  export let totalrecord = 0;
  let dispatch = createEventDispatcher();
  let listKeluaran = [];
  let listPrediksi = [];
  let title_page = "PASARAN";
  let sData = "";
  let myModal_newentry = "";
  let pasaran_nama = "";
  let pasaran_url = "";
  let pasaran_diundi = "";
  let pasaran_jamopen = "";
  let pasaran_display = 0;
  let pasaran_slug = "";
  let pasaran_meta_title = "";
  let pasaran_meta_descp = "";
  let pasaran_status = "";
  let pasaran_create = "";
  let pasaran_update = "";
  let tanggal_keluaran = "";
  let nomor_keluaran = "";
  let tanggal_prediksi = "";
  let bbfs_prediksi = "";
  let nomor_prediksi = "";
  let idrecord = "";
  let record = "";
  let record_message = "";
  let totalrecord_keluaran = "";
  let css_loader = "display: none;";
  let msgloader = "";
  let info_bbfs = ""
  let info_prediksi = ""
  let info_jituaskop = ""
  let info_jitukepalaekor = ""
  let info_dasar = ""
  let info_shio = ""
  let info_2d = ""
  let info_3d = ""
  
  const RefreshHalaman = () => {
    dispatch("handleRefreshData", "call");
  };
  const crudData = (e,id,nama, url, diundi, jamopen, display, status, slug, meta_title,meta_descp,create,update) => {
    sData = e;
    if(sData == "New"){
      clearField();
    }else{
      idrecord = id;
      pasaran_nama = nama;
      pasaran_url = url;
      pasaran_diundi = diundi;
      pasaran_jamopen = jamopen;
      pasaran_display = display;
      pasaran_slug = slug;
      pasaran_meta_title = meta_title;
      pasaran_meta_descp = meta_descp;
      pasaran_create = create;
      pasaran_update = update;
      if (status == "SHOW") {
        status = "Y";
      }
      pasaran_status = status;
    }
    myModal_newentry = new bootstrap.Modal(document.getElementById("modalentry"));
    myModal_newentry.show();
  }
  const EditData = (e, nama, url, diundi, jamopen, display, status) => {
    sData = "Edit";
    
    myModal_newentry = new bootstrap.Modal(
      document.getElementById("modalentryedit")
    );
    myModal_newentry.show();
  };
  const RemoveKeluaran = (e) => {
    handleDeleteKeluaran(e);
  };
  const RemovePrediksi = (e) => {
    handleDeletePrediksi(e);
  };
  const ShowKeluaran = (e, nama) => {
    sData = "Edit";
    idrecord = e;
    pasaran_nama = nama;
    myModal_newentry = new bootstrap.Modal(
      document.getElementById("modallistkeluaran")
    );
    myModal_newentry.show();
    call_keluaran();
  };
  const ShowPrediksi = (e, nama) => {
    sData = "Edit";
    idrecord = e;
    pasaran_nama = nama;
    myModal_newentry = new bootstrap.Modal(
      document.getElementById("modallistprediksi")
    );
    myModal_newentry.show();
    call_prediksi();
  };
  const ShowInfoPrediksi = (nama,prediksi) => {
    pasaran_nama = nama;
    let temp = prediksi.split(" - ")
    let temp_nomor = temp[2]
    let temp_dasar = parseInt(temp_nomor[1]) + parseInt(temp_nomor[2])
    let temp_dasar_final = ""
    if(temp_dasar > 9){
      let temp_dasar2 = parseInt(temp_dasar[1]) + parseInt(temp_dasar[2])
      temp_dasar = temp_dasar2
    }
    let data_1 = ""
    let data_2 = ""
    if(temp_dasar <= 4) {
			data_1 = "KECIL"
		} else {
			data_1 = "BESAR"
		}
		if(temp_dasar%2 == 0) {
			data_2 = "GENAP"
		} else {
			data_2 = "GANJIL"
		}
    temp_dasar_final = data_1+" - "+data_2
    info_bbfs = temp[1]
    info_prediksi = temp[2]
    info_jituaskop = temp_nomor[0]+" dan "+temp_nomor[1]
    info_jitukepalaekor = temp_nomor[1]+" dan "+temp_nomor[2]
    info_dasar = temp_dasar_final
    info_shio = temp_nomor[1]+""+temp_nomor[2]

    let temp_array = [];
    let temp2_array = [];
    let temp3 = "";
    let temp4 = "";
    for (let i = 0; i < info_bbfs.length; i++) {
      temp_array.push(info_bbfs[i]);
    }
    temp2_array = shuffle(temp_array);
    for (let i = 0; i < info_bbfs.length; i++) {
      temp3 += temp2_array[i];
    }
    let xc = 0
    for (let i = 0; i < temp3.length; i++) {
      for (let j = 0; j < temp3.length; j++) {
        if(temp3[i] != temp3[j]){
          xc = xc + 1;
          if(xc <= 20){
            if(xc == 20){
              temp4 += temp3[i]+""+temp3[j];
            }else{
              temp4 += temp3[i]+""+temp3[j]+"*";
            }
          }
        }
      }
    }
    info_2d = temp4;
    temp4 = "";
    xc = 0;
    
    for (let i = 0; i < temp3.length; i++) {
      for (let j = 0; j < temp3.length; j++) {
        for (let z = 0; z < temp3.length; z++) {
          let flag = true
          if(temp3[i] == temp3[j]){
            flag = false           
          }
          if(temp3[i] == temp3[z]){
            flag = false         
          }
          if(temp3[j] == temp3[z]){
            flag = false         
          }
          if(flag){
            xc = xc + 1;
            if(xc <= 20){
              if(xc == 20){
                temp4 += temp3[i]+""+temp3[j]+""+temp3[z];
              }else{
                temp4 += temp3[i]+""+temp3[j]+""+temp3[z]+"*";
              }
            }
          }
        }
      }
    }
    info_3d = temp4;
    myModal_newentry = new bootstrap.Modal(
      document.getElementById("modalinfoprediksi")
    );
    myModal_newentry.show();
  };
  const ShowNewKeluaran = () => {
    myModal_newentry = new bootstrap.Modal(
      document.getElementById("modalentrykeluaran")
    );
    myModal_newentry.show();
  };
  const ShowNewPrediksi = () => {
    myModal_newentry = new bootstrap.Modal(
      document.getElementById("modalentryprediksi")
    );
    myModal_newentry.show();
  };
  async function handleSave() {
    let flag = true;
    let msg = "";
    if (sData == "New") {
      if (idrecord == "") {
        flag = false;
        msg += "The Code is required\n";
      }
      if (pasaran_nama == "") {
        flag = false;
        msg += "The Pasaran is required\n";
      }
      if (pasaran_url == "") {
        flag = false;
        msg += "The URL is required\n";
      }
      if (pasaran_diundi == "") {
        flag = false;
        msg += "The Hari diundi is required\n";
      }
      if (pasaran_jamopen == "") {
        flag = false;
        msg += "The Jam open is required\n";
      }
    } else {
      if (pasaran_nama == "") {
        flag = false;
        msg += "The Pasaran is required\n";
      }
      if (pasaran_url == "") {
        flag = false;
        msg += "The URL is required\n";
      }
      if (pasaran_diundi == "") {
        flag = false;
        msg += "The Hari diundi is required\n";
      }
      if (pasaran_jamopen == "") {
        flag = false;
        msg += "The Jam open is required\n";
      }
    }

    if (flag) {
      css_loader = "display: inline-block;";
      msgloader = "Sending...";
      const res = await fetch("/api/pasaransave", {
        method: "POST",
        headers: {
          "Content-Type": "application/json",
          Authorization: "Bearer " + token,
        },
        body: JSON.stringify({
          sdata: sData,
          page: "KELUARAN-SAVE",
          pasaran_id: idrecord,
          pasaran_name: pasaran_nama,
          pasaran_url: pasaran_url,
          pasaran_diundi: pasaran_diundi,
          pasaran_jamjadwal: pasaran_jamopen,
          pasaran_display: parseInt(pasaran_display),
          pasaran_status: pasaran_status,
          pasaran_slug: pasaran_slug,
          pasaran_meta_title: pasaran_meta_title,
          pasaran_meta_descp: pasaran_meta_descp,
        }),
      });
      const json = await res.json();
      if (json.status == 200) {
        clearField();
        msgloader = json.message;
        myModal_newentry.hide();
        RefreshHalaman();
      } else if (json.status == 403) {
        alert(json.message);
      } else {
        msgloader = json.message;
      }
      setTimeout(function () {
        css_loader = "display: none;";
      }, 1000);
    } else {
      alert(msg);
    }
  }
  async function handleSaveKeluaran() {
    let flag = true;
    let msg = "";
    if (tanggal_keluaran == "") {
      flag = false;
      msg += "The Tanggal is required\n";
    }
    if (nomor_keluaran == "") {
      flag = false;
      msg += "The Nomor is required\n";
    }
    if (flag) {
      css_loader = "display: inline-block;";
      msgloader = "Sending...";
      const res = await fetch("/api/keluaransave", {
        method: "POST",
        headers: {
          "Content-Type": "application/json",
          Authorization: "Bearer " + token,
        },
        body: JSON.stringify({
          sdata: sData,
          page: "KELUARAN-SAVE",
          pasaran_id: idrecord,
          keluaran_tanggal: tanggal_keluaran,
          keluaran_nomor: nomor_keluaran,
        }),
      });
      const json = await res.json();
      if (json.status == 200) {
        call_keluaran();
        RefreshHalaman();
        nomor_keluaran = "";
        msgloader = json.message;
      } else if (json.status == 403) {
        alert(json.message);
      } else {
        msgloader = json.message;
      }
      setTimeout(function () {
        css_loader = "display: none;";
      }, 1000);
    } else {
      alert(msg);
    }
  }
  async function handleDeleteKeluaran(e) {
    let flag = true;
    let msg = "";
    if (e == "") {
      flag = false;
      msg = "The Keluaran is required";
    }
    if (flag) {
      css_loader = "display: inline-block;";
      msgloader = "Sending...";
      const res = await fetch("/api/keluarandelete", {
        method: "POST",
        headers: {
          "Content-Type": "application/json",
          Authorization: "Bearer " + token,
        },
        body: JSON.stringify({
          page: "KELUARAN-DELETE",
          pasaran_id: idrecord,
          keluaran_id: e,
        }),
      });
      const json = await res.json();
      if (json.status == 200) {
        call_keluaran();
        RefreshHalaman();
        msgloader = json.message;
      } else if (json.status == 403) {
        alert(json.message);
      } else {
        msgloader = json.message;
      }
      setTimeout(function () {
        css_loader = "display: none;";
      }, 1000);
    } else {
      alert(msg);
    }
  }
  async function handleSavePrediksi() {
    let flag = true;
    let msg = "";
    if (tanggal_prediksi == "") {
      flag = false;
      msg += "The Tanggal is required\n";
    }
    if (bbfs_prediksi == "") {
      flag = false;
      msg += "The BBFS is required\n";
    }
    if (nomor_prediksi == "") {
      flag = false;
      msg += "The Nomor is required\n";
    }
    if (flag) {
      css_loader = "display: inline-block;";
      msgloader = "Sending...";
      const res = await fetch("/api/prediksisave", {
        method: "POST",
        headers: {
          "Content-Type": "application/json",
          Authorization: "Bearer " + token,
        },
        body: JSON.stringify({
          sdata: sData,
          page: "KELUARAN-SAVE",
          pasaran_id: idrecord,
          prediksi_tanggal: tanggal_prediksi,
          Prediksi_bbfs: bbfs_prediksi,
          Prediksi_nomor: nomor_prediksi,
        }),
      });
      const json = await res.json();
      if (json.status == 200) {
        call_prediksi();
        RefreshHalaman();
        nomor_prediksi = "";
        bbfs_prediksi = "";
        msgloader = json.message;
      } else if (json.status == 403) {
        alert(json.message);
      } else {
        msgloader = json.message;
      }
      setTimeout(function () {
        css_loader = "display: none;";
      }, 1000);
    } else {
      alert(msg);
    }
  }
  async function handleDeletePrediksi(e) {
    let flag = true;
    let msg = "";
    if (e == "") {
      flag = false;
      msg = "The Prediksi is required";
    }
    if (flag) {
      css_loader = "display: inline-block;";
      msgloader = "Sending...";
      const res = await fetch("/api/prediksidelete", {
        method: "POST",
        headers: {
          "Content-Type": "application/json",
          Authorization: "Bearer " + token,
        },
        body: JSON.stringify({
          page: "KELUARAN-DELETE",
          pasaran_id: idrecord,
          prediksi_id: e,
        }),
      });
      const json = await res.json();
      if (json.status == 200) {
        call_prediksi();
        RefreshHalaman();
        msgloader = json.message;
      } else if (json.status == 403) {
        alert(json.message);
      } else {
        msgloader = json.message;
      }
      setTimeout(function () {
        css_loader = "display: none;";
      }, 1000);
    } else {
      alert(msg);
    }
  }
  async function call_keluaran() {
    listKeluaran = [];
    css_loader = "display: inline-block;";
    msgloader = "Fetching...";
    const res = await fetch("/api/keluaran", {
      method: "POST",
      headers: {
        "Content-Type": "application/json",
        Authorization: "Bearer " + token,
      },
      body: JSON.stringify({
        page: "KELUARAN-view",
        pasaran_id: idrecord,
      }),
    });
    const json = await res.json();
    if (json.status == 200) {
      record = json.record;
      record_message = json.message;
      if (record != null) {
        totalrecord_keluaran = record.length;
        for (var i = 0; i < record.length; i++) {
          listKeluaran = [
            ...listKeluaran,
            {
              keluaran_id: record[i]["keluaran_id"],
              keluaran_tanggal: record[i]["keluaran_tanggal"],
              keluaran_periode: record[i]["keluaran_periode"],
              Keluaran_nomor: record[i]["keluaran_nomor"],
            },
          ];
        }
      }
    }
    msgloader = json.message;
    setTimeout(function () {
      css_loader = "display: none;";
    }, 1000);
  }
  async function call_prediksi() {
    listPrediksi = [];
    css_loader = "display: inline-block;";
    msgloader = "Fetching...";
    const res = await fetch("/api/prediksi", {
      method: "POST",
      headers: {
        "Content-Type": "application/json",
        Authorization: "Bearer " + token,
      },
      body: JSON.stringify({
        page: "KELUARAN-VIEW",
        pasaran_id: idrecord,
      }),
    });
    const json = await res.json();
    if (json.status == 200) {
      record = json.record;
      record_message = json.message;
      if (record != null) {
        totalrecord_keluaran = record.length;
        for (var i = 0; i < record.length; i++) {
          listPrediksi = [
            ...listPrediksi,
            {
              prediksi_id: record[i]["prediksi_id"],
              prediksi_tanggal: record[i]["prediksi_tanggal"],
              prediksi_bbfs: record[i]["prediksi_bbfs"],
              prediksi_nomor: record[i]["prediksi_nomor"],
            },
          ];
        }
      }
    }
    msgloader = json.message;
    setTimeout(function () {
      css_loader = "display: none;";
    }, 1000);
  }
  function clearField() {
    idrecord = "";
    pasaran_nama = "";
    pasaran_diundi = "";
    pasaran_jamopen = "";
    pasaran_display = 0;
    pasaran_status = "";
    pasaran_slug = "";
    pasaran_meta_title = "";
    pasaran_meta_descp = "";
    pasaran_create = "";
    pasaran_update = "";
  }
  function callFunction(event) {
    switch (event.detail) {
      case "NEW":
        crudData("New","","","","","","","","","","","","");
        break;
      case "NEW_KELUARAN":
        ShowNewKeluaran();
        break;
      case "SAVE_KELUARAN":
        handleSaveKeluaran();
        break;
      case "NEW_PREDIKSI":
        ShowNewPrediksi();
        break;
      case "SAVE_PREDIKSI":
        handleSavePrediksi();
        break;
      case "REFRESH":
        RefreshHalaman();
        break;
      case "SAVE":
        handleSubmit();
        break;
    }
  }
  const handleKeyboard_formatstring = () => {
    let text = idrecord;
    idrecord = text.toUpperCase();
    let temp = "";
    for (let i = 0; i < idrecord.length; i++) {
      temp = idrecord.charCodeAt(idrecord[i]);
      if (parseInt(temp) < 65 || parseInt(temp) > 90) {
        idrecord = "";
      }
    }
  };
  const handleKeyboard_format = () => {
    let numbera;

    for (let i = 0; i < pasaran_jamopen.length; i++) {
      numbera = parseInt(pasaran_jamopen[i]);
      if (isNaN(numbera)) {
        if (pasaran_jamopen[i] != ":") {
          pasaran_jamopen = "";
        }
      }
    }
    for (let i = 0; i < pasaran_display.length; i++) {
      numbera = parseInt(pasaran_display[i]);
      if (isNaN(numbera)) {
        pasaran_display = "";
      }
    }
    for (let i = 0; i < nomor_keluaran.length; i++) {
      numbera = parseInt(nomor_keluaran[i]);
      if (isNaN(numbera)) {
        nomor_keluaran = "";
      }
    }
    for (let i = 0; i < nomor_prediksi.length; i++) {
      numbera = parseInt(nomor_prediksi[i]);
      if (isNaN(numbera)) {
        nomor_prediksi = "";
      }
    }
    for (let i = 0; i < bbfs_prediksi.length; i++) {
      numbera = parseInt(bbfs_prediksi[i]);
      if (isNaN(numbera)) {
        bbfs_prediksi = "";
      }
    }
  };
  function generate() {
    let numbergenerate = (Math.floor(Math.random() * 1000000) + 1000000)
      .toString()
      .substring(1);
    let temp_array = [];
    let temp2_array = [];
    let temp3 = "";
    bbfs_prediksi = numbergenerate;
    for (let i = 0; i < bbfs_prediksi.length; i++) {
      temp_array.push(bbfs_prediksi[i]);
    }
    temp2_array = shuffle(temp_array);
    for (let i = 0; i < 4; i++) {
      temp3 += temp2_array[i];
    }
    nomor_prediksi = temp3;
  }
  function shuffle(array) {
    return array.sort(() => Math.random() - 0.5);
  }
  function convert_time(e) {
    let temp = e.split(":")
    let hour = temp[0]
    let minute = temp[1]
    return hour+":"+minute
  }
</script>

<div id="loader" style="margin-left:50%;{css_loader}">
  {msgloader}
</div>
<div class="container" style="margin-top: 70px;">
  <div class="row">
    <div class="col-sm-12">
      <Button
        on:click={callFunction}
        button_function="NEW"
        button_title="New"
        button_css="btn-dark"/>
      <Button
        on:click={callFunction}
        button_function="REFRESH"
        button_title="Refresh"
        button_css="btn-primary"/>
      <Panel card_title={title_page} card_footer={totalrecord}>
        <slot:template slot="card-body">
          <table class="table table-striped ">
            <thead>
              <tr>
                <th NOWRAP width="1%" style="text-align: center;vertical-align: top;" colspan="4">&nbsp;</th>
                <th NOWRAP width="1%" style="text-align: center;vertical-align: top;font-weight:bold;font-size:{table_header_font};">NO</th>
                <th NOWRAP width="1%" style="text-align: center;vertical-align: top;font-weight:bold;font-size:{table_header_font};">STATUS</th>
                <th NOWRAP width="1%" style="text-align: left;vertical-align: top;font-weight:bold;font-size:{table_header_font};">CODE</th>
                <th NOWRAP width="*" style="text-align: left;vertical-align: top;font-weight:bold;font-size: {table_header_font};">PASARAN</th>
                <th NOWRAP width="20%" style="text-align: left;vertical-align: top;font-weight:bold;font-size: {table_header_font};">HARI DIUNDI</th>
                <th NOWRAP width="10%" style="text-align: center;vertical-align: top;font-weight:bold;font-size: {table_header_font};">JADWAL</th>
                <th NOWRAP width="10%" style="text-align: left;vertical-align: top;font-weight:bold;font-size: {table_header_font};">KELUARAN</th>
                <th NOWRAP width="10%" style="text-align: left;vertical-align: top;font-weight:bold;font-size: {table_header_font};">PREDIKSI</th>
              </tr>
            </thead>
            {#if totalrecord > 0}
              <tbody>
                {#each listHome as rec}
                  <tr>
                    <td NOWRAP style="text-align: center;vertical-align: top;cursor:pointer;">
                      <i on:click={() => {
                          crudData("Edit",
                            rec.pasaran_id,
                            rec.pasaran_name,
                            rec.pasaran_url,
                            rec.pasaran_diundi,
                            convert_time(rec.pasaran_jamjadwal),
                            rec.pasaran_display,
                            rec.pasaran_status,
                            rec.pasaran_slug,
                            rec.pasaran_meta_title,
                            rec.pasaran_meta_descp,
                            rec.pasaran_create,
                            rec.pasaran_update
                          );
                        }}
                        class="bi bi-pencil"/></td>
                    <td NOWRAP style="text-align: center;vertical-align: top;cursor:pointer;">
                      <i on:click={() => {
                          ShowKeluaran(rec.pasaran_id, rec.pasaran_name);
                        }}
                        class="bi bi-card-list"/>
                    </td>
                    <td NOWRAP style="text-align: center;vertical-align: top;cursor:pointer;">
                      <i
                        on:click={() => {
                          ShowPrediksi(rec.pasaran_id, rec.pasaran_name);
                        }}
                        class="bi bi-file-binary"/>
                    </td>
                    <td NOWRAP style="text-align: center;vertical-align: top;cursor:pointer;">
                      <i
                        on:click={() => {
                          ShowInfoPrediksi(rec.pasaran_name,rec.pasaran_prediksi);
                        }}
                        class="bi bi-book"/>
                    </td>
                    <td NOWRAP style="text-align: center;vertical-align: top;font-size: {table_body_font};">{rec.pasaran_no}</td>
                    <td NOWRAP style="text-align: center;vertical-align: top;font-size: {table_body_font};">
                      <span style="padding: 5px;border-radius: 10px;padding-right:10px;padding-left:10px;{rec.pasaran_statuscss}">
                        {rec.pasaran_status}
                      </span>
                    </td>
                    <td NOWRAP style="text-align: left;vertical-align: top;font-size: {table_body_font};">{rec.pasaran_id}</td>
                    <td NOWRAP style="text-align: left;vertical-align: top;font-size: {table_body_font};">
                      <a href={rec.pasaran_url} target="_blank">{rec.pasaran_name}</a>
                    </td>
                    <td NOWRAP style="text-align: left;vertical-align: top;font-size: {table_body_font};">{rec.pasaran_diundi}</td>
                    <td NOWRAP style="text-align: center;vertical-align: top;font-size: {table_body_font};">{convert_time(rec.pasaran_jamjadwal)}</td>
                    <td NOWRAP style="text-align: left;vertical-align: top;font-size: {table_body_font};">{rec.pasaran_keluaran}</td>
                    <td NOWRAP style="text-align: left;vertical-align: top;font-size: {table_body_font};">{rec.pasaran_prediksi}</td>
                  </tr>
                {/each}
              </tbody>
            {:else}
              <tbody>
                <tr>
                  <td colspan="20">
                    <center>
                      <Loader />
                    </center>
                  </td>
                </tr>
              </tbody>
            {/if}
          </table>
        </slot:template>
      </Panel>
    </div>
  </div>
</div>

<Modal
  modal_id="modalentry"
  modal_size="modal-dialog-centered modal-lg"
  modal_title={title_page + "/" + sData}
  modal_footer_css="padding:5px;"
  modal_footer={true}>
  <slot:template slot="body">
    <div class="row">
      <div class="col-sm-6">
        <div class="mb-3">
          <label for="exampleForm" class="form-label">Code</label>
          <Input
            bind:value={idrecord}
            on:keyup={handleKeyboard_formatstring}
            class="required"
            minlength="2"
            maxlength="4"
            type="text"
            placeholder="Code"/>
        </div>
        <div class="mb-3">
          <label for="exampleForm" class="form-label">Pasaran</label>
          <Input
            bind:value={pasaran_nama}
            class="required"
            type="text"
            placeholder="Pasaran"/>
        </div>
        <div class="mb-3">
          <label for="exampleForm" class="form-label">URL</label>
          <Input
            bind:value={pasaran_url}
            class="required"
            type="text"
            placeholder="URL"/>
        </div>
        <div class="mb-3">
          <label for="exampleForm" class="form-label">Hari diundi</label>
          <Input
            bind:value={pasaran_diundi}
            class="required"
            type="text"
            placeholder="Hari diundi"/>
        </div>
        <div class="mb-3">
          <label for="exampleForm" class="form-label">JADWAL</label>
          <Input
            bind:value={pasaran_jamopen}
            on:keyup={handleKeyboard_format}
            class="required"
            style="text-align:center;"
            type="text"
            placeholder="Jadwal"/>
        </div>
        <div class="mb-3">
          <label for="exampleForm" class="form-label">Display</label>
          <Input
            bind:value={pasaran_display}
            on:keyup={handleKeyboard_format}
            class="required"
            type="text"
            style="text-align:right;"
            placeholder="Display"/>
        </div>
        <div class="mb-3">
          <label for="exampleForm" class="form-label">Status</label>
          <select class="form-control required" bind:value={pasaran_status}>
            <option value="Y">SHOW</option>
          </select>
        </div>
      </div>
      <div class="col-sm-6">
        <div class="mb-3">
          <label for="exampleForm" class="form-label">Slug</label>
          <Input
            bind:value={pasaran_slug}
            class="required"
            type="text"
            placeholder="Slug"/>
            <p class="text-muted">Ex : keluaran-togel-hk or keluaran-togel-cambodia</p>
        </div>
        <div class="mb-3">
          <label for="exampleForm" class="form-label">Meta - Title</label>
          <Input
            bind:value={pasaran_meta_title}
            class="required"
            type="text"
            maxlength=60
            placeholder="Meta - Title"/>
        </div>
        <div class="mb-3">
          <label for="exampleForm" class="form-label">Meta - Deskripsi</label>
          <textarea
            style="height: 100px;resize: none;"
            bind:value={pasaran_meta_descp}
            class="form-control required"/>
        </div>
        {#if sData == "Edit"}
            <div class="alert alert-dark" role="alert" style="font-size:11px;padding:5px;">
                Create : {pasaran_create}
                <br />
                Update : {pasaran_update}
            </div>
        {/if}
        <div class="alert alert-info" role="alert" style="font-size:12px;padding:5px;">
          Ex :<br /> 
          Title SEO : <br />Keluaran Togel Cambodia - Keluaran NOMOR TOGEL
          <br /> <br />
          Descp SEO : <br />Result Togel Cambodia, Result nomor togel, keluaran nomor togel, result cambodia
        </div>
      </div>
    </div>
  </slot:template>
  <slot:template slot="footer">
    <Button
      on:click={() => {
        handleSave();
      }}
      button_function="SAVE"
      button_title="Save"
      button_css="btn-warning"/>
  </slot:template>
</Modal>



<Modal
  modal_id="modallistkeluaran"
  modal_size="modal-dialog-centered"
  modal_title="Keluaran : {pasaran_nama}"
  modal_body_css="padding:5px;margin:0px;height:500px;overflow-y: scroll;"
  modal_footer_css="padding:5px;"
  modal_footer={true}>
  <slot:template slot="body">
    <table class="table">
      <thead>
        <tr>
          <th NOWRAP width="1%" style="text-align: center;vertical-align: top;font-weight:bold;font-size:{table_header_font};">&nbsp;</th>
          <th NOWRAP width="20%" style="text-align: center;vertical-align: top;font-weight:bold;font-size:{table_header_font};">DATE</th>
          <th NOWRAP width="10%" style="text-align: center;vertical-align: top;font-weight:bold;font-size:{table_header_font};">PERIODE</th>
          <th NOWRAP width="*" style="text-align: left;vertical-align: top;font-weight:bold;font-size:{table_header_font};">NOMOR</th>
        </tr>
      </thead>
      <tbody>
        {#each listKeluaran as rec}
          <tr>
            <td NOWRAP style="text-align: center;vertical-align: top;cursor:pointer;">
              <i on:click={() => {
                  RemoveKeluaran(rec.keluaran_id);
                }}
                class="bi bi-trash"/>
            </td>
            <td NOWRAP style="text-align: center;vertical-align: top;font-size: {table_body_font};">{dayjs(rec.keluaran_tanggal).format("YYYY-MM-DD")}</td>
            <td NOWRAP style="text-align: center;vertical-align: top;font-size: {table_body_font};">{rec.keluaran_periode}</td>
            <td NOWRAP style="text-align: left;vertical-align: top;font-size: {table_body_font};">{rec.Keluaran_nomor}</td>
          </tr>
        {/each}
      </tbody>
    </table>
  </slot:template>
  <slot:template slot="footer">
    <Button
      on:click={callFunction}
      button_function="NEW_KELUARAN"
      button_title="New"
      button_css="btn-primary"/>
  </slot:template>
</Modal>

<Modal
  modal_id="modalentrykeluaran"
  modal_size="modal-dialog-centered"
  modal_title="Keluaran {pasaran_nama + '/' + sData}"
  modal_footer_css="padding:5px;"
  modal_footer={true}>
  <slot:template slot="body">
    <div class="mb-3">
      <label for="exampleForm" class="form-label">Tanggal</label>
      <Input
        bind:value={tanggal_keluaran}
        class="required"
        type="date"
        name="date"
        id="exampleDate"
        data-date-format="dd-mm-yyyy"
        placeholder="date placeholder"/>
    </div>
    <div class="mb-3">
      <label for="exampleForm" class="form-label">Nomor</label>
      <Input
        bind:value={nomor_keluaran}
        on:keyup={handleKeyboard_format}
        class="required"
        minlength="4"
        maxlength="4"
        type="text"
        placeholder="Nomor Keluaran"/>
    </div>
  </slot:template>
  <slot:template slot="footer">
    <Button
      on:click={callFunction}
      button_function="SAVE_KELUARAN"
      button_title="Save"
      button_css="btn-warning"/>
  </slot:template>
</Modal>

<Modal
  modal_id="modallistprediksi"
  modal_size="modal-dialog-centered"
  modal_title="Prediksi : {pasaran_nama}"
  modal_body_css="padding:5px;margin:0px;height:500px;overflow-y: scroll;"
  modal_footer_css="padding:5px;"
  modal_footer={true}>
  <slot:template slot="body">
    <table class="table">
      <thead>
        <tr>
          <th NOWRAP width="1%" style="text-align: center;vertical-align: top;font-weight:bold;font-size:{table_header_font};">&nbsp;</th>
          <th NOWRAP width="*" style="text-align: center;vertical-align: top;font-weight:bold;font-size:{table_header_font};">DATE</th>
          <th NOWRAP width="20%" style="text-align: center;vertical-align: top;font-weight:bold;font-size:{table_header_font};">BBFS</th>
          <th NOWRAP width="20%" style="text-align: center;vertical-align: top;font-weight:bold;font-size:{table_header_font};">NOMOR</th>
        </tr>
      </thead>
      <tbody>
        {#each listPrediksi as rec}
          <tr>
            <td NOWRAP style="text-align: center;vertical-align: top;cursor:pointer;">
              <i on:click={() => {
                  RemovePrediksi(rec.prediksi_id);
                }}
                class="bi bi-trash"/>
            </td>
            <td NOWRAP style="text-align: center;vertical-align: top;font-size: {table_body_font};">{dayjs(rec.prediksi_tanggal).format("YYYY-MM-DD")}</td>
            <td NOWRAP style="text-align: center;vertical-align: top;font-size: {table_body_font};">{rec.prediksi_bbfs}</td>
            <td NOWRAP style="text-align: center;vertical-align: top;font-size: {table_body_font};">{rec.prediksi_nomor}</td>
          </tr>
        {/each}
      </tbody>
    </table>
  </slot:template>
  <slot:template slot="footer">
    <Button
      on:click={callFunction}
      button_function="NEW_PREDIKSI"
      button_title="New"
      button_css="btn-primary"/>
  </slot:template>
</Modal>

<Modal
  modal_id="modalentryprediksi"
  modal_size="modal-dialog-centered"
  modal_title="Prediksi {pasaran_nama + '/' + sData}"
  modal_footer_css="padding:5px;"
  modal_footer={true}>
  <slot:template slot="body">
    <div class="mb-3">
      <label for="exampleForm" class="form-label">Tanggal</label>
      <Input
        bind:value={tanggal_prediksi}
        class="required"
        type="date"
        name="date"
        id="exampleDate"
        data-date-format="dd-mm-yyyy"
        placeholder="date placeholder"/>
    </div>
    <div class="mb-3">
      <label for="exampleForm" class="form-label">BBFS</label>
      <div class="input-group mb-3">
        <Input
          bind:value={bbfs_prediksi}
          on:keyup={handleKeyboard_format}
          class="required"
          minlength="6"
          maxlength="6"
          type="text"
          placeholder="Nomor BBFS"/>
        <button
          on:click={() => {
            generate();
          }}
          type="button"
          class="btn btn-info">Generate</button>
      </div>
    </div>
    <div class="mb-3">
      <label for="exampleForm" class="form-label">Nomor</label>
      <Input
        bind:value={nomor_prediksi}
        on:keyup={handleKeyboard_format}
        class="required"
        minlength="6"
        maxlength="6"
        type="text"
        placeholder="Nomor Prediksi"/>
    </div>
  </slot:template>
  <slot:template slot="footer">
    <Button
      on:click={callFunction}
      button_function="SAVE_PREDIKSI"
      button_title="Save"
      button_css="btn-warning"/>
  </slot:template>
</Modal>

<Modal
  modal_id="modalinfoprediksi"
  modal_size="modal-dialog-centered"
  modal_title="Prediksi : {pasaran_nama}"
  modal_body_css="padding:5px;margin:0px;height:400px;overflow-y: scroll;"
  modal_footer_css="padding:5px;"
  modal_footer={false}>
  <slot:template slot="body">
    <p>
      BBFS: {info_bbfs}<br />
      4D {pasaran_nama}: {info_prediksi}<br />
      COLOK BEBAS {pasaran_nama}: {info_prediksi}<br />
      COLOK JITU AS DAN KOP {pasaran_nama}: {info_jituaskop}<br />
      COLOK JITU KEPALA DAN EKOR {pasaran_nama}: {info_jitukepalaekor}<br />
      DASAR : {info_dasar}<br />
      SHIO : {info_shio}<br />
      20 LINE 3D JITU : {info_3d}<br />
      20 LINE 2D JITU : {info_2d}<br />
    </p>
  </slot:template>
  <slot:template slot="footer">
    
  </slot:template>
</Modal>