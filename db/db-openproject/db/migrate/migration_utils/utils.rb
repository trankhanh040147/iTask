#-- copyright
# OpenProject is an open source project management software.
# Copyright (C) 2012-2024 the OpenProject GmbH
#
# This program is free software; you can redistribute it and/or
# modify it under the terms of the GNU General Public License version 3.
#
# OpenProject is a fork of ChiliProject, which is a fork of Redmine. The copyright follows:
# Copyright (C) 2006-2013 Jean-Philippe Lang
# Copyright (C) 2010-2013 the ChiliProject Team
#
# This program is free software; you can redistribute it and/or
# modify it under the terms of the GNU General Public License
# as published by the Free Software Foundation; either version 2
# of the License, or (at your option) any later version.
#
# This program is distributed in the hope that it will be useful,
# but WITHOUT ANY WARRANTY; without even the implied warranty of
# MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
# GNU General Public License for more details.
#
# You should have received a copy of the GNU General Public License
# along with this program; if not, write to the Free Software
# Foundation, Inc., 51 Franklin Street, Fifth Floor, Boston, MA  02110-1301, USA.
#
# See COPYRIGHT and LICENSE files for more details.
#++

module Migration
  module Utils
    UpdateResult = Struct.new(:row, :updated)

    def say_with_time_silently(message, &)
      say_with_time message do
        suppress_messages(&)
      end
    end

    def in_configurable_batches(klass, default_batch_size: 1000)
      batches = ENV["OPENPROJECT_MIGRATION_BATCH_SIZE"]&.to_i || default_batch_size

      klass.in_batches(of: batches)
    end

    def remove_index_if_exists(table_name, index_name)
      if index_name_exists? table_name, index_name
        remove_index table_name, name: index_name
      end
    end

    ##
    # Executes the given SQL query while passing in sanitized parameters.
    #
    # @param query [String] SQL query including parameter references like `:param`
    # @param params [Hash] Hash containing values for referenced parameters
    #
    # @raise [ActiveRecord::ActiveRecordError] If the query fails
    # @return [PG::Result]
    #
    # Example:
    #
    #   execute_sql "select id from users where mail = :email", email: params[:email]
    #
    def execute_sql(query, params = {})
      query = ActiveRecord::Base.sanitize_sql [query, params]

      ActiveRecord::Base.connection.execute query
    end
  end
end
