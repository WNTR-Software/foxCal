import { Component, signal } from "@angular/core";
import { RouterOutlet } from "@angular/router";
import { BaseLayout } from "~/layouts/base/base.layout";

@Component({
	selector: "app-root",
	imports: [RouterOutlet, BaseLayout],
	template: `
		<app-base>
			<h1
				class="text-3xl lg:text-5xl mb-6 font-bold underline font-title"
			>
				Welcome to {{ title() }}!
			</h1>

			<p class="text-primary">Primary</p>
			<p class="text-secondary">Secondary</p>
			<p class="text-accent">Accent</p>
			<pre class="font-mono">Code Example</pre>

			<router-outlet />
		</app-base>
	`,
	styles: [],
})
export class App {
	protected readonly title = signal("foxCal");
}
