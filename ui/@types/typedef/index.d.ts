declare module "typedef" {
    export default namespace td {
        type ProblemStruct = { title: string, link: string, description: string, };
        type DiscussionEntry = { title: string, link: string, };
        type DiscussionStruct = { alt: string, links: DiscussionEntry[], };
        type ResponseStruct = { error: null | string, content: ProblemStruct | DiscussionStruct, };
    }
}
