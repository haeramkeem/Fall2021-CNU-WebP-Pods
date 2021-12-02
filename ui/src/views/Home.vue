// Main page
//      @path: /

<template>
    <div>
        <div class="container">
            <div class="row justify-content-center align-items-center main-row">
                <mini-calendar :minify="false" @dateChanged="onDateChanged"/>
            </div>
            <div class="row main-row">
                <problem-box :problem="this.problem"/>
                <discussion-box @langChanged="onLanguageChanged" :discussion="this.discussion"/>
            </div>
        </div>
    </div>
</template>

<script lang="ts">
import { defineComponent } from 'vue';
import td from 'typedef';
import axios from 'axios';

import MiniCalendar from '../components/MiniCalendar.vue';
import ProblemBox from '../components/ProblemBox.vue';
import DiscussionBox from '../components/DiscussionBox.vue';

export default defineComponent({
    name: "Home",
    components: {
        MiniCalendar,
        ProblemBox,
        DiscussionBox,
    },
    data() {
        return {
            problem: {} as td.ProblemStruct,
            discussion: {} as td.DiscussionStruct,
            selectedLang: "all",
        };
    },
    methods: {
        /**
         * GET problem for selected date
         * @params date: string
         */
        onDateChanged(date: string): void {
            this.discussion = {};
            this.problem = {};
            axios.get("/api/main/problem?date=" + date)
            .then((resp) => {
                const data = resp.data as td.ResponseStruct;
                if(typeof data.error === "string") {
                    throw new Error(data.error);
                } else {
                    this.problem = data.content as td.ProblemStruct;
                    this.fetchDiscussion(this.selectedLang, this.parseProblemPath());
                }
            }).catch((err) => {
                alert(err);
            });
        },
        /**
         * GET discussion when language changed
         * @params lang: string
         */
        onLanguageChanged(lang: string): void {
            this.selectedLang = lang;
            this.fetchDiscussion(lang, this.parseProblemPath());
        },
        /**
         * GET discussion by selected language & problem
         * @params lang: string, problemPath: string
         */
        fetchDiscussion(lang:string, problemPath: string): void {
            this.discussion = {};
            axios.get(`/api/discussion/${problemPath}?lang=${lang}`)
            .then((resp) => {
                const data = resp.data as td.ResponseStruct;
                if(typeof data.error === "string") {
                    throw new Error(data.error);
                } else {
                    this.discussion = data.content as td.DiscussionStruct;
                }
            }).catch((err) => {
                alert(err);
            });
        },
        /**
         * Parse problem path from problem link
         * @return string
         */
        parseProblemPath(): string {
            const urlPart = this.problem.link.split("/");
            const len = urlPart.length;
            return urlPart[len - 1].length === 0 ? urlPart[len - 2] : urlPart[len - 1];
        }
    },
});
</script>

<style scoped>
.main-row {
    min-height: 192px;
}
</style>
