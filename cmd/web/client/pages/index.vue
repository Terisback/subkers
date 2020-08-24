<template>
    <div class="container">
        <div class="header">
            <img class="logo" src="/subkers.png" alt="Subkers Logo" />
            <h1 class="title">Subkers</h1>
        </div>
        <div class="application">
            <p class="info">Tool for converting subtitles into markers for Adobe Audition</p>
            <div class="convert-container">
                <button class="convert-button" @click.prevent="$refs.file.click()">
                    <span>Upload</span>
                </button>
            </div>
            <input
                type="file"
                ref="file"
                accept=".ass, .ssa, .srt, .stl, .ttml, .vtt"
                style="display: none"
                @change="onChange"
            />
        </div>
        <div class="footer">
            <span>Made by</span>
            <a href="https://terisback.ru/">Anton Terisback</a>
        </div>
    </div>
</template>

<script>
import { saveAs } from "file-saver";
import axios from "axios";

//var fileTypes = ["srt", "ssa", "ass", "stl", "ttml", "vtt"];

async function change(event) {
    var file = event.target.files[0];
    var reader = new FileReader();
    reader.onload = async function (e) {
        var formData = new FormData();
        var newFilename = file.name.split(".")
        formData.append("extension", newFilename.pop());
        var sendBlob = new Blob([e.target.result]);
        formData.append("subtitle", sendBlob, file.name);
        axios
            .post("/api/convert", formData, {
                headers: {
                    "Content-Type": "multipart/form-data",
                },
            })
            .then((responce) => {
                var blob = new Blob([responce.data], {
                    type: "text/plain;charset=utf-8",
                });
                saveAs(blob, newFilename + ".csv");
            });
    };

    reader.readAsText(file);
}

export default {
    methods: {
        onChange: change,
    },
};
</script>

<style>
.container {
    margin: 0 auto;
    max-width: 350px;
    height: 100vh;
    display: flex;
    justify-content: center;
    align-items: center;
    text-align: center;
    flex-direction: column;
    gap: 16px;
}

.header {
    margin-top: auto;
    width: 100%;
    height: 100px;
    display: flex;
    flex-direction: row;
    justify-content: center;
    align-content: center;
    text-align: center;
}

.title {
    margin-left: 16px;
    font-family: "system";
    font-weight: 500;
    font-size: 48px;
    line-height: 100px;
    letter-spacing: 1px;
}

.logo {
    height: 100%;
}

.application {
    width: 100%;
    padding: 24px;
    background-color: #405780;
    border-radius: 8px;
    display: flex;
    flex-direction: column;
    justify-content: center;
    align-content: center;
    text-align: center;
    gap: 16px;
}

.convert-container {
    border-radius: 12px;
    overflow: hidden;
    background-image: linear-gradient(145deg, #1897ff, #0072d3);
}

.convert-container:hover {
    background-image: linear-gradient(145deg, #19a3ff, #117bd3);
}

.convert-button {
    width: 100%;
    height: 100%;
    cursor: pointer;
    border: 0;
    padding: 16px;
    background-color: transparent;
}

.convert-button:focus {
    outline: 0;
}

.convert-button span {
    font-family: "system";
    font-size: 48px;
    font-weight: 600;
    letter-spacing: 8px;
    text-align: center;
}

.footer {
    margin-top: auto;
    margin-bottom: 10px;
}

.footer span,
.footer a {
    color: #84aef5;
}
</style>
