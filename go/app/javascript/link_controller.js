export default class LinkController extends Stimulus.Controller {

  static values = { href: String }

  async delete() {
    console.log(`DELETE ${this.hrefValue}`)
    const response = await fetch(this.hrefValue, { method: 'DELETE' })

    console.log(response)
    if (response.redirected) {
      location.replace(response.url)
    }
  }
}
