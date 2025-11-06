import { Component } from "@angular/core";
import { SiteHeader } from "~/components/core/site-header/site-header";

@Component({
	selector: "app-base",
	imports: [SiteHeader],
	template: `
		<app-site-header />
		<main class="container mx-auto p-4 lg:px-12 bg-base text-base-content">
			<ng-content />
		</main>
	`,
	styles: ``,
})
export class BaseLayout {}
