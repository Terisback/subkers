<template>
    <div class="container">
        <div class="header">
            <img class="logo" src="/subkers.png" alt="Subkers Logo" />
            <h1 class="title">Subkers</h1>
        </div>
        <div class="application">
            <p class="info">Tool for converting subtitles into markers for Adobe Audition</p>
            <button class="convert-button">
                <span>Upload</span>
            </button>
        </div>
        <div class="footer">
            <span>Made by</span>
            <a href="https://terisback.ru/">Anton Terisback</a>
        </div>
    </div>
</template>

<script>
export default {};
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

.convert-button {
    cursor: pointer;
    border: 0;
    padding: 16px;
    border-radius: 12px;
    transition: 1s;
    background: linear-gradient(145deg, #1897ff, #0072d3);
}

.convert-button:hover {
    background: linear-gradient(145deg, #19a3ff, #117bd3);
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

.footer{
    margin-top: auto;
    margin-bottom: 10px;
}

.footer span, .footer a{
    color:#587fc4
}
</style>
