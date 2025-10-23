import { bootstrapApplication } from '@angular/platform-browser';
import { appConfig } from './app/app.config';
import { App } from './app/app';
import "@fontsource/atkinson-hyperlegible"
import "@fontsource/atkinson-hyperlegible-mono"
import "@fontsource/nunito-sans"

bootstrapApplication(App, appConfig).catch((err) => console.error(err));
