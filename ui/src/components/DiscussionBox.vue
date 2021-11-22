<template>
    <div class="col no-pad">
        <div class="container no-pad">
            <h2 class="row justify-content-between align-items-center">
                <div class="no-pad col-auto fira-sans font-bold">Discussions</div>
                <div class="no-pad col-auto lang-select">
                    <span class="fira-sans font-bold">LANGUAGE</span>
                    <select @change="onLanguageChanged($event.target.value)"> 
                        <option value="all">All</option>
                        <option value="cpp">C++</option>
                        <option value="java">Java</option>
                        <option value="python">Python</option>
                    </select>
                </div>
            </h2>
        </div>
        <article>
            <div v-if="this.discussion.alt" class="fetch-nothing">
                <h3 class="bebas-neue"><i class="bi bi-emoji-neutral-fill"></i> Damn! We can't fetch any discussions for this problem.</h3>
                <a :href="this.discussion.alt" target="_blank" class="alt-link bebas-neue">Show me althernative link instead..</a>
            </div>
            <div v-else>
                <ul>
                    <li v-for="(item, idx) in this.discussion.links" v-bind:key="idx">
                        <a :href="item.link" class="disc-link" target="_blank">{{item.title}}</a>
                    </li>
                </ul>
            </div>
        </article>
    </div>
</template>

<script lang="ts">
import { defineComponent, PropType } from 'vue';
import td from 'typedef';

export default defineComponent({
    name: "DiscussionBox",
    props: {
        discussion: Object as PropType<td.DiscussionStruct>,
    },
    methods: {
        onLanguageChanged(lang: string): void {
            this.$emit("langChanged", lang);
        }
    },
    created() {
        this.$emit("langChanged", "all");
    },
});
</script>

<style scoped>
h2 {
    font-size: 28px;
    border-bottom: 1px solid black;
    margin: 0px 15px 8px 15px;
}

.lang-select {
    font-size: 18px;
    border: 2px solid #285CBF;
    height: 24px;
    padding-right: 15px;
}

.lang-select>span {
    background-color: #285CBF;
    color: white;
    padding: 0px 15px;
    margin-right: 15px;
}

.lang-select>select {
    all: unset;
}

article {
    padding: 0px 15px;
}

.fetch-nothing {
    text-align: center;
    color: #FF5733;
}

.alt-link {
    font-size: 24px;
}

ul {
    list-style: none;
    padding-left: 0px;
}

li {
    padding: 9px 0px;
    border-bottom: 1px solid lightgray;
}

.disc-link {
    color: black;
    text-decoration: none;
    font-size: 18px;
}

.disc-link:hover {
    text-decoration: underline;
}
</style>
