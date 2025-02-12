{
  "name": "{{ .PackageNS }}-client-ts",
  "version": "0.0.1",
  "description": "Autogenerated Typescript Client",
  "author": "Ignite Codegen <hello@ignite.com>",
  "license": "Apache-2.0",
  "licenses": [
    {
      "type": "Apache-2.0",
      "url": "http://www.apache.org/licenses/LICENSE-2.0"
    }
  ],
  "main": "lib/index.js",
  "publishConfig": {
    "access": "public"
  },
  "scripts": {
    "build": "NODE_OPTIONS='--max-old-space-size=16384' tsc",
    "postinstall": "npm run build"
  },
  "dependencies": {    
    "@cosmjs/proto-signing": "0.31.1",
    "@cosmjs/stargate": "0.31.1",
    "@keplr-wallet/types": "^0.11.3", 
    "axios": "0.21.4",
    "buffer": "^6.0.3",
    "events": "^3.3.0"
  },
  "peerDependencies": {
    "@cosmjs/proto-signing": "0.31.1",
    "@cosmjs/stargate": "0.31.1"
  }, 
  "devDependencies": {
    "@types/events": "^3.0.0",
    "typescript": "^4.8.4"
  }
}
