<template>
    <div>
        <header class="container-fluid">
            <div class="row px-5">
                <h1 class="col-12 col-lg">
                    <a href="/" id="logo">
                        <span class="hidden">pods</span>
                    </a>
                </h1>
                <nav class="col-12 col-lg">
                    <router-link to="/difficulty" class="open-sans font-bold router-link">
                        by difficulty
                    </router-link>
                    <span> | </span>
                    <router-link to="/topic" class="open-sans font-bold router-link">
                        by topic
                    </router-link>
                </nav>
            </div>
        </header>
        <main>
            <router-view></router-view>
        </main>
        <footer class="fira-sans font-bold">
            copyright©saltwalks2021
        </footer>
    </div>
</template>

<style scoped>
* {
    margin: 0px;
    padding: 0px;
}
.row {
    min-height: 140px;
    background-color: black;
}

#logo {
    display: inline-block;
    width: 270px;
    height: 70px;
    background-image: url("./assets/img/logo.png");
    background-repeat: no-repeat;
    background-size: contain;
}

h1 {
    height: 140px;
    padding: 35px 0px;
}

nav {
    padding: 45px 0px;
    text-align: right;
    font-size: 24px;
    color: white;
}

.router-link {
    display: inline-block;
    padding: 7px 20px;
    text-decoration: none;
    text-transform: uppercase;
    color: white;
}

.router-link:hover {
    color: white;
    background-color: #285CBF;
}

@media(max-width: 992px) {
    h1 {
        padding: 10px 0px;
        text-align: center;
        height: 90px;
    }
    nav {
        padding: 10px 0px;
        text-align: center;
        font-size: 24px;
        color: white;
    }
}

footer {
    display: block;
    width: 100%;
    text-align: center;
    background-color: lightgray;
    position: fixed;
    top: calc(100% - 20px);
}
</style>
