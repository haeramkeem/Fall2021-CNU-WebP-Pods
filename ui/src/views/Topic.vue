// Page for problem topic
//      @path: /topic

<template>
    <div>
        <div class="container">
            <div class="row flex-column justify-content-center align-items-center main-row">
                <mini-calendar :minify="true"/>
                <radio-box :items="this.radioItems" @radioClicked="onRadioClicked"/>
            </div>
            <div class="row main-row">
                <problem-box :problem="this.problem" :titlePrefix="this.radioItems[this.selectedTopic]"/>
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
import RadioBox from '../components/RadioBox.vue';
import ProblemBox from '../components/ProblemBox.vue';
import DiscussionBox from '../components/DiscussionBox.vue';

export default defineComponent({
    name: "Topic",
    components: {
        MiniCalendar,
        RadioBox,
        ProblemBox,
        DiscussionBox,
    },
    data() {
        return {
            radioItems: ["bfs", "dfs", "dp", "two pointer", "sort", "greedy", "binary search", "tree", "heap", "string"],
            selectedTopic: 0,
            problem: {} as td.ProblemStruct,
            discussion: {} as td.DiscussionStruct,
            selectedLang: "all",
        };
    },
    methods: {
        /**
         * GET problem for selected topic
         * @params idx: number
         */
        onRadioClicked(idx: number): void {
            axios.get("/api/topic/problem?topic=" + idx)
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
            this.selectedTopic = idx;
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
    created(): void {
        // Init page with first radio item clicked
        this.onRadioClicked(0);
    }
});
</script>

<style scoped>
.main-row {
    min-height: 192px;
}
</style>
