FROM ruby:2.7.1

# Add node repo
RUN curl -sL https://deb.nodesource.com/setup_lts.x | bash -

# Add yarn repo
RUN curl -sS https://dl.yarnpkg.com/debian/pubkey.gpg | apt-key add -
RUN echo "deb https://dl.yarnpkg.com/debian/ stable main" | tee /etc/apt/sources.list.d/yarn.list

RUN apt-get update -qq && apt-get install -y \
    nodejs \
    yarn \
    postgresql-client \
    && rm -rf /var/lib/apt/lists/*

WORKDIR /app

COPY Gemfile Gemfile.lock ./
RUN bundle check || bundle install

COPY package.json yarn.lock ./
RUN yarn install --check-files

ENV RAILS_SERVE_STATIC_FILES true
ENV RAILS_LOG_TO_STDOUT true

CMD ["rails", "server", "-b", "0.0.0.0"]
