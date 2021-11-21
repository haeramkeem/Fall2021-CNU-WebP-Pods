declare module "typedef" {
    export default namespace td {
        type ProblemStruct = { title: string, link: string, description: string, };
        type ResponseStruct = { error: null | string, content: ProblemStruct, };
    }
}
