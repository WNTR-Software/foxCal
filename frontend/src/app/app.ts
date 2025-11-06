import { Component } from "@angular/core";
import { RouterOutlet } from "@angular/router";
import { BaseLayout } from "~/layouts/base/base.layout";

@Component({
	selector: "app-root",
	imports: [RouterOutlet, BaseLayout],
	template: `
		<app-base>
			<router-outlet />
		</app-base>
	`,
	styles: [],
})
export class App {
}
