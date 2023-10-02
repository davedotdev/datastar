package webui

import (
	"bytes"
	"compress/gzip"
	"context"
	"fmt"
	"net/http"

	"github.com/delaneyj/gomponents-iconify/iconify/carbon"
	"github.com/delaneyj/gomponents-iconify/iconify/game_icons"
	"github.com/delaneyj/gomponents-iconify/iconify/gridicons"
	"github.com/delaneyj/gomponents-iconify/iconify/material_symbols"
	"github.com/delaneyj/gomponents-iconify/iconify/mdi"
	"github.com/delaneyj/gomponents-iconify/iconify/ph"
	"github.com/delaneyj/gomponents-iconify/iconify/svg_spinners"
	"github.com/delaneyj/gomponents-iconify/iconify/tabler"
	"github.com/delaneyj/gomponents-iconify/iconify/vscode_icons"
	"github.com/delaneyj/gomponents-iconify/iconify/zondicons"
	. "github.com/delaneyj/toolbelt/gomps"
	"github.com/dustin/go-humanize"
	"github.com/go-chi/chi/v5"
)

func setupHome(ctx context.Context, router *chi.Mux) error {
	build, err := staticFS.ReadFile("static/datastar.iife.js")
	if err != nil {
		return fmt.Errorf("error reading build: %w", err)
	}
	buf := bytes.NewBuffer(nil)
	w, err := gzip.NewWriterLevel(buf, gzip.BestCompression)
	if err != nil {
		return fmt.Errorf("error compressing build: %w", err)
	}

	if _, err := w.Write(build); err != nil {
		return fmt.Errorf("error compressing build: %w", err)
	}
	w.Close()
	iifeBuildSize := humanize.Bytes(uint64(buf.Len()))

	type Feature struct {
		Description string
		Icon        NODE
		Details     NODE
	}

	features := []Feature{
		{
			Description: "Fine Grained Reactivity via Signals",
			Icon:        ph.GitDiff(),
			Details:     DIV(TXT("No Virtual DOM. proxy wrappers, or re-rendering the entire page on every change.  Take the best available options and use hassle free.")),
		},
		{
			Description: "Fully Compliant",
			Icon:        mdi.LanguageHtmlFive(),
			Details:     DIV(TXT("No monkey patching, no custom elements, no custom attributes, no custom anything.  Just plain old HTML5.")),
		},
		{
			Description: "Everything is an Plugin",
			Icon:        gridicons.Plugins(),
			Details:     DIV(TXT("Disagree with the built-in behavior? No problem, just write your own extension in a type safe way.  Take what you need, leave what you don't.")),
		},
		{
			Description: "Batteries Included (but optional)",
			Icon:        game_icons.Batteries(),
			Details: DIV(
				CLS("breadcrumbs"),
				UL(
					CLS(
						"flex flex-wrap gap-2 justify-center items-center",
					),
					LI(TXT("Custom Actions")),
					LI(TXT("Attribute Binding")),
					LI(TXT("Focus")),
					LI(TXT("Signals")),
					LI(TXT("DOM Events")),
					LI(TXT("Refs")),
					LI(TXT("Intersects")),
					LI(TXT("Two-Way Binding")),
					LI(TXT("Visibility")),
					LI(TXT("Teleporting")),
					LI(TXT("Text Replacement")),
					LI(TXT("HTMX like features")),
					LI(TXT("Server Sent Events")),
				),
			),
		},
	}

	type nodeChildFn func(...NODE) NODE
	languages := []nodeChildFn{
		vscode_icons.FileTypeAssembly,
		vscode_icons.FileTypeApl,
		vscode_icons.FileTypeC,
		vscode_icons.FileTypeCeylon,
		vscode_icons.FileTypeCpp,
		vscode_icons.FileTypeCobol,
		vscode_icons.FileTypeCoffeescript,
		vscode_icons.FileTypeClojure,
		vscode_icons.FileTypeCrystal,
		vscode_icons.FileTypeCsharp,
		vscode_icons.FileTypeDartlang,
		vscode_icons.FileTypeElixir,
		vscode_icons.FileTypeErlang,
		vscode_icons.FileTypeFsharp,
		vscode_icons.FileTypeFortran,
		vscode_icons.FileTypeGoGopher,
		vscode_icons.FileTypeGroovy,
		vscode_icons.FileTypeHaskell,
		vscode_icons.FileTypeHaxe,
		vscode_icons.FileTypeJava,
		vscode_icons.FileTypeJs,
		vscode_icons.FileTypeJulia,
		vscode_icons.FileTypeKotlin,
		vscode_icons.FileTypeLisp,
		vscode_icons.FileTypeLolcode,
		vscode_icons.FileTypeLua,
		vscode_icons.FileTypeNim,
		vscode_icons.FileTypeOcaml,
		vscode_icons.FileTypePerl,
		vscode_icons.FileTypePhp,
		vscode_icons.FileTypePython,
		vscode_icons.FileTypeR,
		vscode_icons.FileTypeRuby,
		vscode_icons.FileTypeRust,
		vscode_icons.FileTypeScala,
		vscode_icons.FileTypeShell,
		vscode_icons.FileTypeSwift,
		vscode_icons.FileTypeTypescript,
		vscode_icons.FileTypeVb,
		vscode_icons.FileTypeZig,
	}

	router.Get("/", func(w http.ResponseWriter, r *http.Request) {
		Render(w, Page(
			DIV(
				CLS("flex-1 flex flex-wrap md:p-16 text-xl flex-col items-center text-center bg-gradient-to-tr from-base-100 to-base-200"),

				DIV(
					DIV(
						H1(
							CLS("text-6xl font-bold"),
							TXT("Declarative HTML "),
						),
						P(
							TXT("Delieved in a blazing fast, simple single framework"),
						),
						DIV(
							CLS("flex gap-2 justify-center items-center"),
						),
					),
					CLS("max-w-4xl flex flex-col items-center justify-center gap-16"),
					DIV(
						CLS("flex flex-wrap gap-6 justify-center items-center"),
						DIV(
							CLS("flex flex-wrap gap-2 justify-center items-center text-6xl"),
							RANGE(languages, func(fn nodeChildFn) NODE {
								return DIV(
									CLS("avatar avatar-xl"),
									fn(
										CLS("w-24 h-24 mask bg-gradient-to-t from-base-200 to-base-300 p-4 mask-hexagon"),
									),
								)
							})),
						A(
							CLS("link-accent text-4xl"),
							HREF("https://htmx.org/essays/hypermedia-on-whatever-youd-like/"),
							TXT("HTML on whatever you want"),
						),
					),
					DIV(
						CLS("flex flex-col gap-2 w-full"),
						H3(
							CLS("text-3xl font-bold"),
							TXT("Simple count example code"),
						),
						DIV(
							CLS("bg-base-100 shadow-inner text-base-content p-4 rounded-box"),
							HIGHLIGHT("html", `<div data-signal-count="0">
	<div>
		<button data-on-click="$count++">Increment +</button>
		<button data-on-click="$count--">Decrement -</button>
		<input type="number" data-model="count" />
	</div>
	<div data-text="$count"></div>
</div>
`,
							),
						),
						DIV(
							CLS("flex gap-2 justify-center items-center"),
							DIV(
								CLS("badge badge-accent flex-1 gap-1"),
								tabler.FileZip(),
								TXT(iifeBuildSize+" w/ all extensions"),
							),
							DIV(
								CLS("badge badge-accent flex-1 gap-1"),
								carbon.ColumnDependency(),
								TXTF("%d Dependencies", len(packageJSON.Dependencies)),
							),
							DIV(
								CLS("badge badge-accent flex-1 gap-1"),
								zondicons.Checkmark(),
								TXT("Fully Tree Shakeable"),
							),
						),
					),
					DIV(
						CLS("flex flex-col gap-2 w-full"),
						H3(
							CLS("text-3xl font-bold"),
							TXT("Global count example from Backend"),
						),
						DIV(
							ID("global-count-example"),
							CLS("flex justify-center p-4 items-center gap-2"),
							DATA("signal-get", "'/api/globalCount'"),
							DATA("on-load", "@get"),
							SPAN(TXT("Loading example on delay...")),
							svg_spinners.Eclipse(
								CLS("datastar-indicator"),
							),
						),
						H5(
							CLS("text-2xl font-bold"),
							TXT("Open the console to see the Fetch/XHR traffic"),
						),
					),
					DIV(
						CLS("card w-full shadow-2xl ring-4 bg-base-300 ring-secondary text-secondary-content"),
						DIV(
							CLS("card-body flex flex-col justify-center items-center"),
							UL(
								CLS("flex flex-col gap-6 justify-center items-center text-2xl gap-4  max-w-xl"),
								RANGE(features, func(f Feature) NODE {
									return LI(
										DIV(
											CLS("flex flex-col gap-1 justify-center items-center"),
											DIV(
												CLS("flex gap-2 items-center"),
												f.Icon,
												TXT(f.Description),
											),
											DIV(
												CLS("text-lg opacity-50 p-2 rounded"),

												f.Details,
											),
										),
									)
								}),
							),
						),
					),

					DIV(
						CLS("flex flex-col gap-2 justify-center items-center"),
						TXT("Built with "),
						DIV(
							CLS("flex gap-1 justify-center items-center text-5xl"),
							vscode_icons.FileTypeHtml(),
							material_symbols.AddRounded(),
							vscode_icons.FileTypeTypescriptOfficial(),
							material_symbols.AddRounded(),
							vscode_icons.FileTypeVite(),
							material_symbols.AddRounded(),
							vscode_icons.FileTypeGoGopher(),
						),
						DIV(
							CLS("flex gap-2 justify-center items-center"),
							TXT("by "),
							A(
								CLS("link-accent"),
								HREF("http://github.com/delaneyj"),
								TXT("Delaney"),
							),
							TXT("and looking for contributors!"),
						),
					),
					DIV(
						CLS("w-full flex gap-2 items-center"),
						A(
							CLS("btn btn-lg flex-1"),
							HREF("/essays/2023-09-08_why-another-framework"),
							material_symbols.Help(),
							TXT("Why another framework?"),
						),
						A(
							CLS("btn btn-primary btn-lg flex-1"),
							HREF("/docs"),
							mdi.RocketLaunch(),
							TXT("Don't care, just get started"),
						),
					),
				),
			),
		),
		)
	})

	return nil
}
