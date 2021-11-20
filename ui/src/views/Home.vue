<template>
    <div>
        <div class="container">
            <div class="row justify-content-center align-items-center main-row">
                <MiniCalendar :minify="false" @dateChanged="onDateChanged"/>
            </div>
            <div class="row main-row">
                <div class="col problem-box no-pad">
                    <h2 class="fira-sans font-bold">Problem of the Day</h2>
                    <article>
                        <a :href="this.problemLink" target="_blank"><h3>{{this.problemTitle}}</h3></a>
                        <div id="fetched-problem"></div>
                    </article>
                </div>
                <div class="col discussion-box no-pad">
                    <div class="container no-pad">
                        <h2 class="row fira-sans font-bold justify-content-between align-items-center">
                            <div class="no-pad col-auto">Discussions</div>
                            <div class="no-pad col-auto" id="lang-select">
                                <span>LANGUAGE</span>
                                <select @change="onLanguageChanged($event.target.value)"> 
                                    <option value="all">All</option>
                                    <option value="cpp">C++</option>
                                    <option value="java">Java</option>
                                    <option value="python">Python</option>
                                </select>
                            </div>
                        </h2>
                    </div>
                    <article>TODO: fetch discussion</article>
                </div>
            </div>
        </div>
    </div>
</template>

<script lang="ts">
import { defineComponent } from 'vue';
import MiniCalendar from '../components/MiniCalendar.vue';
import { td } from 'typedef';
import axios from 'axios';

export default defineComponent({
    name: "Home",
    components: {
        MiniCalendar,
    },
    data() {
        return {
            problemTitle: "",
            problemLink: "",
        };
    },
    methods: {
        onDateChanged(date: Date) {
            console.warn("TODO: fetch problem & discussion for " + date.toString());
        },
        onLanguageChanged(lang: string) {
            console.warn("TODO: language changed to " + lang);
        },
    },
    created() {
        axios.get("/api/main/problem")
        .then((resp) => {
            const refined = resp.data as td.ProblemStruct;
            if(typeof refined.error === "string") {
                throw new Error(refined.error);
            } else {
                this.problemTitle = refined.title;
                this.problemLink = refined.link;
                const target = document.getElementById("fetched-problem") as HTMLElement;
                target.innerHTML = refined.content;
            }
        }).catch((err) => {
            alert(err);
        });
    },
});
</script>

<style scoped>
.main-row {
    min-height: 192px;
}

h2 {
    font-size: 28px;
    border-bottom: 1px solid black;
    margin: 0px 18px;
}
article {
    height: 100%;
}

.problem-box>article {
    border-right: 1px solid black;
    padding: 15px;
}

.discussion-box>h2 {
    text-align: justify;
}

#lang-select {
    font-size: 18px;
    border: 2px solid #285CBF;
    height: 23px;
    padding-right: 15px;
}

#lang-select>span {
    background-color: #285CBF;
    color: white;
    padding: 0px 15px;
    margin-right: 15px;
}

#lang-select>select {
    all: unset;
}
</style>
