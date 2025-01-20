// Code generated by templ - DO NOT EDIT.

// templ: version: v0.2.747
package components

//lint:file-ignore SA4006 This context is only used if a nested component is present.

import "github.com/a-h/templ"
import templruntime "github.com/a-h/templ/runtime"

func AddRecipe() templ.Component {
	return templruntime.GeneratedTemplate(func(templ_7745c5c3_Input templruntime.GeneratedComponentInput) (templ_7745c5c3_Err error) {
		templ_7745c5c3_W, ctx := templ_7745c5c3_Input.Writer, templ_7745c5c3_Input.Context
		templ_7745c5c3_Buffer, templ_7745c5c3_IsBuffer := templruntime.GetBuffer(templ_7745c5c3_W)
		if !templ_7745c5c3_IsBuffer {
			defer func() {
				templ_7745c5c3_BufErr := templruntime.ReleaseBuffer(templ_7745c5c3_Buffer)
				if templ_7745c5c3_Err == nil {
					templ_7745c5c3_Err = templ_7745c5c3_BufErr
				}
			}()
		}
		ctx = templ.InitializeContext(ctx)
		templ_7745c5c3_Var1 := templ.GetChildren(ctx)
		if templ_7745c5c3_Var1 == nil {
			templ_7745c5c3_Var1 = templ.NopComponent
		}
		ctx = templ.ClearChildren(ctx)
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<div id=\"modal-wrapper\" class=\"fixed top-0 left-0 w-full h-full flex items-center justify-center\"><div id=\"modal-underlay\" class=\"absolute top-0 left-0 w-full h-full bg-black bg-opacity-50\" _=\"on click remove #modal-wrapper\"></div><div class=\"modal relative bg-[#fdbf98] w-96 p-8 rounded-lg z-10\"><button class=\"close absolute bg-[#fdbf98] text-white top-0 right-0 p-4 cursor-pointer\" _=\"on click remove #modal-wrapper\"><svg xmlns=\"http://www.w3.org/2000/svg\" class=\"h-6 w-6 text-white\" fill=\"none\" viewBox=\"0 0 24 24\" stroke=\"currentColor\"><path strokeLinecap=\"round\" strokeLinejoin=\"round\" strokeWidth=\"2\" d=\"M6 18L18 6M6 6l12 12\"></path></svg></button><h2 class=\"text-lg font-[900] mb-4\">Add a new recipe</h2><form class=\"bg-[#ffb675] p-4\" hx-post=\"/recipe\" hx-encoding=\"multipart/form-data\" hx-target=\"#recipes-container\" hx-swap=\"innerHTML\" _=\"on htmx:afterRequest remove #modal-wrapper\"><label class=\"block mb-4\">Name: <input type=\"text\" name=\"name\" class=\"block w-full border border-gray-300 rounded-md p-2 bg-white text-black\"></label> <label class=\"block mb-4\">Ingredients: <input type=\"text\" name=\"ingredients\" class=\"block w-full border border-gray-30 font-[500] rounded-md p-2 bg-white text-black\"></label> <label class=\"block mb-4\">Preparation: <input type=\"text\" name=\"preparation\" class=\"block w-full border border-gray-300 rounded-md p-2 bg-white text-black\"></label> <label class=\"block mb-4\">Difficulty: <input type=\"text\" name=\"difficulty\" class=\"block w-full border border-gray-300 rounded-md p-2 bg-white text-black\"></label> <label class=\"block mb-4\">Image: <input type=\"file\" name=\"image\" class=\"block w-full border border-gray-300 rounded-md p-2 bg-white text-black\"></label> <button type=\"submit\" class=\"bg-blue-500 text-white rounded-md font-[500] py-2 px-4 hover:bg-blue-600\">Add a recipe!</button></form></div></div>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		return templ_7745c5c3_Err
	})
}
