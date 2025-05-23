import Web from '../../wats/helpers/web.js';

class Benchmark {
  constructor() {
    this.url = `file://${process.argv[2] || '/tmp/benchmark.html'}`;
    this.web = new Web(this.url);
  }

  async run() {
    await this.web.init();
    await this.web.page.goto(this.url);
    await this.web.waitForText('Benchmark Report')
      .catch(e => {
        console.log('wait for text "Benchmark Report failed"', e);
      });
    const bodyHandle = await this.web.page.$('body > div > div');
    const html = await this.web.page.evaluate(body => body.innerText, bodyHandle);
    await bodyHandle.dispose();
    console.log(html);
  }
}

const main = async () => {
  await new Benchmark().run();
  process.exit(0);
};

main();
