<template>
    <dir>
        <Editor :init="init" v-model="content"></Editor>  
    </dir>
</template>

<script>
    import Editor from '@tinymce/tinymce-vue'
    import tinymce from './tinymce.min.js'
    import './icons/default/icons.min.js'
    import './themes/silver/theme.min.js'
    import './langs/zh_CN'
    //注册插件
    import './plugins/preview/plugin.min.js'
    import './plugins/code/plugin.min.js'
    import './plugins/codesample/plugin.min.js'
    import './plugins/image/plugin.min.js'
    import './plugins/paste/plugin.min.js'
    import './plugins/wordcount/plugin.min.js'
    import './plugins/imagetools/plugin.min.js'
    export default {
        components:{
            Editor
        },
        props:{
            value:{    
                type:String,
                default:'',
            }
        },
        data(){
            return{
                init:{
                    language:'zh_CN',
                    height:"500px",
                    margin:0,
                    padding:0,
                    plugins:"preview code codesample image paste wordcount imagetools",
                    branding:false,
                    // toolbar:['undo redo | formatselect | alignleft alignright alighnjustify', 'preview code codesample image paste wordcount']
                    //上传图片
                    images_upload_handler:async( blobInfo,succFun,failFun)=>{
                        let formdata = new FormData()
                        formdata.append('file',blobInfo.blob(),blobInfo.name())
                        const {data:result} = await this.$http.post('upload',formdata)
                        succFun(result.url)
                    }
                },
                content:this.value,
            }
        },
        watch:{
            value(newV){
                    this.content=newV
            },
            content(newV){
                this.$emit('input',newV)
            }
        },
    }
    
</script>

<style >
@import url('./skins/ui/oxide/skin.min.css');
</style>
