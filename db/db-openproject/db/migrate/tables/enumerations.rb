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

require_relative 'base'

class Tables::Enumerations < Tables::Base
  def self.table(migration)
    create_table migration do |t|
      t.string :name, limit: 30, default: '', null: false
      t.integer :position, default: 1
      t.boolean :is_default, default: false, null: false
      t.string :type
      t.boolean :active, default: true, null: false
      t.integer :project_id
      t.integer :parent_id

      t.index %i[id type], name: 'index_enumerations_on_id_and_type'
      t.index :project_id, name: 'index_enumerations_on_project_id'
    end
  end
end
